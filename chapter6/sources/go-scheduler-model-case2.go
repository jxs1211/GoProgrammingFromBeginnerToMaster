package main

import (
	"fmt"
	"runtime"
	"time"
)

func deadloop() {
	for {
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("NUMCPU: ", runtime.NumCPU())
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}
