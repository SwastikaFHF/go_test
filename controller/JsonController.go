package controller

import (
	"go_test/core"
)

type JsonController struct {
	core.Controller
}

func (this *JsonController) Get() {
	this.TplNames = "view/jsonlint.tpl"
}
