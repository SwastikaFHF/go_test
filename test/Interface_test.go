package main

import (
	"fmt"
	"log"
	// "testing"
)

var mux map[string]do

func init() {
	mux = make(map[string]do)
}

func main() {
	mux["do"] = &MyCon{}
	if v, ok := mux["do"]; ok {
		v.WriteString("ss")
	} else {
		log.Println("2222")
		fmt.Println("11111111")
	}
}

type MyCon struct {
	Controller
}

type Controller struct {
}

type do interface {
	WriteString(str string)
}

func (ctx *Controller) WriteString(str string) {
	log.Println(str)
	fmt.Println("33333333")
}
