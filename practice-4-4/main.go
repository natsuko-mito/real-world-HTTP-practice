package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"query": {"Hello world"},
	}
	res, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	log.Println("Status: ", res.Status)
	log.Println("StatusCode: ", res.StatusCode)
	log.Println("HeaderFields: ", res.Header)
	log.Println("HeaderField Content-Type: ", res.Header.Get("Content-Type"))
}
