package controller

import (
	"go_personal_blog/core"
)

type MainController struct {
	core.Controller
}

func (this *MainController) Get() {
	this.TplNames = "view/index.tpl"
}
