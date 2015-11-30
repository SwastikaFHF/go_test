package controller

import (
	"go_test/core"
)

type MainController struct {
	core.Controller
}

func (this *MainController) Get() {
	this.TplNames = "../go_test/view/index.tpl"
	// this.Ctx.WriteString("hello world")
}
