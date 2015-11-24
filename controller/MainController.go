package controller

import (
	"io/ioutil"
	"strings"
)

type MainController struct {
	Controller
}

func (this *MainController) Get() {
	url := this.Ctx.Req.URL.String()
	if strings.EqualFold("/", url) {
		_, err := ioutil.ReadFile("./static/index.html")
		if err != nil {
			this.Ctx.WriteString("404")
			return
		}
		// this.Ctx.WriteString(contents)
	} else {
		this.Ctx.WriteString("404 not found:\n\n")
	}
}

func (this *MainController) Post() {

}
