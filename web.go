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
	core.SetStaticPath("/css", "static/css")
	core.SetStaticPath("/font", "static/font")
	core.SetStaticPath("/icon", "static/icon")
	core.SetStaticPath("/js", "static/js")
	core.Run()
}
