package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func Test_Tread(t *testing.T) {
	_, err := ioutil.ReadFile("/static/index.html")
	if err != nil {
		log.Println("error!!!")
	} else {
		log.Println("success")
	}
}

func main() {
	_, err := ioutil.ReadFile("../go_test/static/index.html")
	if err != nil {
		log.Println("error!!!")
	} else {
		log.Println("success")
	}
}
