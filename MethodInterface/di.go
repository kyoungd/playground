package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url string
)

// Dependency Inversion
// This is the interface for HttpClient.Get().
// Instead of calling it directly, we'll use an interface to call it.  We can substitue this interface in Unit testing.
type HttpClient interface {
	Get(string) (*http.Response, error)
}

func init() {
	flag.StringVar(&url, "url", "http://google.com", "Which URL do we want to parse?")
	flag.Parse()
}

func main() {
	client := &http.Client{}
	err := send(client, url)

	if err != nil {
		panic(err)
	}
}

// The function send uses the interface HttpClient instead of the http.Client.
// It enables the client.Get to be substituted in unit testing.
func send(client HttpClient, link string) error {
	response, err := client.Get(link)
	if err != nil {
		return err
	}
	if response == nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
