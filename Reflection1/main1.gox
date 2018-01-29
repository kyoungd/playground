package main

import (
	"fmt"
	"io"
	"fmt"
)

func main() {
	c, _ := net.Dial("tcp", ":2100")
	var r io.Reader
	r = c	// now stores (value:c, type descriptor: net.Conn)
	// that's why we can also do this
	if w, ok := r.(io.Writer); ok {
		fmt.Println("We didn't forget there is a writer inside value c")
		w.Write("take a look")
	}
}
