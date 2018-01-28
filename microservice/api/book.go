package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	// define book
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

func (b Book) ToJSON() []byte {
	toJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return toJSON
}

func FromJSON(data []byte) Book {
	var fromJSON Book
	err := json.Unmarshal(data, &fromJSON)
	if err != nil {
		panic(err)
	}
	return fromJSON
}

var books = map[string]Book{
	"0345391802": Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	"0000000000": Book{Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		b := allBooks()
		WriteJSON(w, b)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusOK)
	}
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case http.MethodGet:
		isbn := r.URL.Path[len("/api/books/"):]
		book, found := GetBook(isbn)
		if found {
			WriteJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		existed := UpdateBook(book)
		if existed {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		isbn := r.URL.Path[len("/api/books/"):]
		existed := DeleteBook(isbn)
		if existed {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method"))
	}
}

func allBooks() []Book {
	bs := make([]Book, len(books))
	ix := 0
	for _, b := range books {
		bs[ix] = b
		ix++
	}
	return bs
}

func CreateBook(book Book) (string, bool) {
	_, exist := books[book.ISBN]
	if !exist {
		books[book.ISBN] = book
		return book.ISBN, true
	} else {
		return "", false
	}
}

func WriteJSON(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func GetBook(isbn string) (Book, bool) {
	v, exist := books[isbn]
	if exist {
		return v, true
	} else {
		return Book{}, false
	}
}

func UpdateBook(b Book) bool {
	_, exist := books[b.ISBN]
	if exist {
		books[b.ISBN] = b
	}
	return exist
}

func DeleteBook(isbn string) bool {
	_, exist := books[isbn]
	if exist {
		delete(books, isbn)
	}
	return exist
}
