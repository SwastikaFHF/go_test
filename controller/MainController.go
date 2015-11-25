package controller

import (
	"go_test/core"
	"io/ioutil"
	"log"
	"net/http"
)

type MainController struct {
	core.Controller
}

func (*MainController) Get(rep http.ResponseWriter, res *http.Request) {
	content, err := ioutil.ReadFile("../go_test/static/index.html")
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		rep.Write(content)
	}
}
