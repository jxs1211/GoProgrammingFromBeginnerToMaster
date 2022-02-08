package main

import (
	"fmt"
	"runtime"
	"time"
)

func add(a, b int) int {
	return a + b
}

type I interface {
	M1()
}

type I2 interface {
	M1()
}

func dummy() {
	add(3, 5)
}

func deadloop() {
	for {
		dummy()
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}
