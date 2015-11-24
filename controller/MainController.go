package controller

import (
	"io/ioutil"
	"log"
	// "strings"
)

type MainController struct {
	Controller
}

func (this *MainController) Get() {
	log.Println("get !!!")
	_, err := ioutil.ReadFile("./static/index.html")
	if err != nil {
		this.Ctx.WriteString("404")
		return
	}
}

func (this *MainController) Post() {

}
