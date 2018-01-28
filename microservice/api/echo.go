package api

import (
	"fmt"
	"net/http"
)

func EchoHandlerFunc(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	message := q["message"][0]
	w.Header().Add("content-type", "text/plain")
	fmt.Fprintf(w, message)
}
