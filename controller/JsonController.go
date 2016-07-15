package controller

import (
	"go_personal_blog/core"
)

type JsonController struct {
	core.Controller
}

func (this *JsonController) Get() {
	this.TplNames = "view/jsonlint.tpl"
}
