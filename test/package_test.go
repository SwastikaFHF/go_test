package test

import (
	"fmt"
	// "strings"
	"log"
	"testing"
)

func TestDomain(t *testing.T) {
	var stu Student
	stu.name = "www"
	log.Println(stu)
	stu.Say()
}

type Person struct {
}

type Student struct {
	Person
	name string
}

func (p Person) Say() {
	fmt.Println("p say")
}
