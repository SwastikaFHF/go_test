package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &MyHandler{},
		ReadTimeout: 5 * time.Second,
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = sayHello
	mux["/bye"] = sayBye
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type MyHandler struct {
}

func (*MyHandler) ServeHTTP(rep http.ResponseWriter, res *http.Request) {
	if h, ok := mux[res.URL.String()]; ok {
		h(rep, res)
		return
	}
	io.WriteString(rep, "Main, this is version 3")
}

func sayHello(rep http.ResponseWriter, res *http.Request) {
	log.Println("sayHello")
	io.WriteString(rep, "Hello World, this is version 3")
}

func sayBye(rep http.ResponseWriter, res *http.Request) {
	log.Println("sayBye")
	io.WriteString(rep, "Bye Bye, this is version 3")
}
