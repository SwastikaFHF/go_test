package main

import (
	"go_test/controller"
	"go_test/controller/common"
	"go_test/core"
)

func main() {
	core.Router("/", &controller.MainController{})
	core.Router("/m", &controller.MarkDown{})
	core.Router("/markdown", &controller.MarkDown{})
	core.Router("/welcome", &controller.Welcome{})
	core.Router("/json1", &common.Common{})
	json := &controller.JsonController{}
	core.Router("/json", json)
	core.Router("/j", json)

	core.SetStaticPath("/css", "static/css")
	core.SetStaticPath("/font", "static/font")
	core.SetStaticPath("/icon", "static/icon")
	core.SetStaticPath("/js", "static/js")
	core.SetStaticPath("/images", "static/img")
	core.Run()
}
