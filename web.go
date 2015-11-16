package main

import (
	"go_test/controller"
	"net/http"
	"os"
)

const (
	HostVar = "VCAP_APP_HOST"
	PortVar = "VCAP_APP_PORT"
)

func main() {
	http.HandleFunc("/", index)
	var port string
	if port = os.Getenv(PortVar); port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}

}

func index(res http.ResponseWriter, req *http.Request) {
	controller.Index(res, req)
}
