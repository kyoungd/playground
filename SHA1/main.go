package main

import (
	"crypto/sha1"
	"fmt"
)

func SHA1EncodedText(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func main() {
	s := "sha1 this string"
	fmt.Println(s)
	fmt.Printf("%x\n", SHA1EncodedText(s))
}
