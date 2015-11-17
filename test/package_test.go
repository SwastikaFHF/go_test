package test

import (
	"fmt"
	"testing"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

func TestDomain(t *testing.T) {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:2], c)
	go sum(a[2:4], c)
	go sum(a[4:], c)

	var x int = <-c
	var y int = <-c
	var z int = <-c
	// x, y := <-c, <-c // 从 c 中获取

	fmt.Println(x, y, z)
}
