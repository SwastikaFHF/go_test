package main

import (
	"go_test/controller"
	// "go_test/controller/common"
)

func main() {
	controller.Router("/", &controller.MainController{})
	controller.Run()
}
