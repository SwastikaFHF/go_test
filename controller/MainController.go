package controller

import (
	"go_test/core"
)

type MainController struct {
	core.Controller
}

func (this *MainController) Get() {
	this.TplNames = "view/index.tpl"
}
