package main

import (
	// "go_test/controller"
	// "go_test/controller/common"
	// "log"
	// "net/http"
	// "os"
	"fmt"
	// "runtime"
)

// const (
// 	HostVar = "VCAP_APP_HOST"
// 	PortVar = "VCAP_APP_PORT"
// )

func main() {
	// http.HandleFunc("/", controller.Index)
	// http.HandleFunc("/getimage", common.GetLoadImage)
	// var port string
	// if port = os.Getenv(PortVar); port == "" {
	// 	port = "8080"
	// }
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	panic(err)
	// }
	// log.Fatal(...)
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched()
		fmt.Println(s)
	}
}
