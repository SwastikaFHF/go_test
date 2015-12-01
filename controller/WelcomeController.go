package controller

import (
	"go_test/core"
)

type Welcome struct {
	core.Controller
}

func (this *Welcome) Get() {
	this.TplNames = "../go_test/view/welcome.tpl"
	this.Data["Website"] = "Website"
	this.Data["Email"] = "Email"
}
