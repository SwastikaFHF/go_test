package main

import (
	"go_test/controller"
	"go_test/controller/common"
	"go_test/core"
)

func main() {
	core.Router("/", &controller.MainController{})
	core.Router("/welcome", &controller.Welcome{})
	core.Router("/json", &common.Common{})
	// core.Router("/img/favicon.ico", &common.Image{})
	// core.SetStaticPath("/css", "../go_test/static/css")
	// core.SetStaticPath("/font", "../go_test/static/font")
	// core.SetStaticPath("/icon", "../go_test/static/icon")
	// core.SetStaticPath("/js", "../go_test/static/js")
	core.Run()
}
