package controller

import (
	"go_test/core"
)

type MainController struct {
	core.Controller
}

func (this *MainController) Get() {
	// this.TplNames = "../go_test/static/index.html"
	this.Ctx.WriteString("hello world")
}
