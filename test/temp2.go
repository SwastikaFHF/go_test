package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &MyHandler{})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type MyHandler struct {
}

func (*MyHandler) ServeHTTP(rep http.ResponseWriter, res *http.Request) {
	io.WriteString(rep, "Hello World, this is version 2")
}
