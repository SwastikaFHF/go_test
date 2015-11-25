package main

import (
	"go_test/controller"
	"go_test/core"
)

func main() {
	core.Router("/", &controller.MainController{})
	core.Run()
}
