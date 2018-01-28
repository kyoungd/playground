package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	json := string(book.ToJSON())
	assert.Equal(t, `{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`,
		json, "Book JSON marshalling is incorrect.")
}

func TestFromJSON(t *testing.T) {
	book2 := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	json := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`)
	book1 := FromJSON(json)
	assert.Equal(t, book1, book2, "Book JSON unmarshalling is incorrect.")
}
