package main

import (
	"go_test/controller"
	"go_test/controller/common"
	"net/http"
	"os"
)

const (
	HostVar = "VCAP_APP_HOST"
	PortVar = "VCAP_APP_PORT"
)

func main() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/getimage", common.GetLoadImage)
	http.HandleFunc("/html", controller.GetFormView)
	http.HandleFunc("/htmlvalue", controller.GetFormValue)
	var port string
	if port = os.Getenv(PortVar); port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
