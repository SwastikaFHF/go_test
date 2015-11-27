package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "text/template"
)

func GetFormView(res http.ResponseWriter, req *http.Request) {
	contents, err := ioutil.ReadFile("./static/htmltest.html")
	if err != nil {
		fmt.Fprintf(res, "404")
		return
	}
	fmt.Fprintf(res, "%s\n", contents)
}

func GetFormValue(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	go fmt.Println(req.Form)
	go log.Println(req.Form)
	// fmt.Fprintf(res, "%s\n", req.Form)
	if b, err := json.Marshal(req.Form); err == nil {
		fmt.Fprintf(res, "%s\n", string(b))
	}
	fmt.Fprintf(res, "%s\n", "error")
}
