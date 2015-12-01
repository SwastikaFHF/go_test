package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(rep http.ResponseWriter, res *http.Request) {
	io.WriteString(rep, "Hello World, this is version 1")
}
