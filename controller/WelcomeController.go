package controller

import (
	"go_personal_blog/core"
)

type Welcome struct {
	core.Controller
}

func (this *Welcome) Get() {
	this.TplNames = "view/welcome.tpl"
	this.Data["Website"] = "Website"
	this.Data["Email"] = "Email"
}
