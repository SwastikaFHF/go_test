package test

import (
	"log"
	"testing"
	"time"
)

func TestTread(t *testing.T) {
	go test()
	n := 0
	for n == 100 {
		time.Sleep(1 * time.Second)
		log.Println("TestTread===")
	}
}

func test() {
	n := 0
	for n == 100 {
		time.Sleep(1 * time.Second)
		log.Println("test===")
	}
}
