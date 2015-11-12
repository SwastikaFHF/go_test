package main

import (
	"encoding/json"
	"fmt"
	"go_test/controller/common"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	HostVar = "VCAP_APP_HOST"
	PortVar = "VCAP_APP_PORT"
)

type Person struct {
	Name string `json:"name"`
}

func test() {

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/env", env)
	http.HandleFunc("/ee", commonHandler)
	var port string
	if port = os.Getenv(PortVar); port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}

}

func commonHandler(res http.ResponseWriter, req *http.Request) {

	var per Person
	str := `{"Name":"seee"}`
	json.Unmarshal([]byte(str), &per)
	fmt.Fprint(res, per)
	common.DoPost(res, req)
}

func index(res http.ResponseWriter, req *http.Request) {
	url := req.URL.String()

	if strings.EqualFold("/", url) {
		contents, err := ioutil.ReadFile("./static/index.html")
		if err != nil {
			fmt.Fprintf(res, "404")
			return
		}
		fmt.Fprintf(res, "%s\n", contents)
	} else {
		fmt.Fprint(res, "404 not found:\n\n")
	}

}

func env(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "System Environment:\n\n")
	env := os.Environ()
	for _, e := range env {
		fmt.Fprintln(res, e)
	}

}
