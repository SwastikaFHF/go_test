package controller

import (
	"go_test/core"
	"os"
	"path"
)

type MainController struct {
	core.Controller
}

func (this *MainController) Get() {
	file, err := os.Getwd()
	if err == nil {
		this.TplNames = path.Join(file, "view/index.tpl")
	}
}
