package controller

import (
	"go_test/core"
)

type MarkDown struct {
	core.Controller
}

func (this *MarkDown) Get() {
	this.TplNames = "view/markdown.tpl"
	this.Data["Website"] = "Website"
	this.Data["Email"] = "Email"
}
