package main

import (
	"fmt"
	"runtime"
	"time"
)

func deadloop() {
	for {
		// time.Sleep(time.Second * 1)
		// fmt.Println("deadloop I got scheduled!")
	}
}

func main() {
	// runtime.GOMAXPROCS(1)
	fmt.Println("NUMCPU: ", runtime.NumCPU())
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}
