package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	sync.Mutex
	i int
}

var cter counter

func Increase() int {
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func showChanCase5() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			v := Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
		}(i)
	}

	time.Sleep(5 * time.Second)
}

type Counter struct {
	i int
	sync.Mutex
}

var counter2 Counter

func Increase2() int {
	counter2.Lock()
	defer counter2.Unlock()
	counter2.i++
	return counter2.i
}

func showChanCase52() {
	for i := 0; i < 10; i++ {
		go func(id int) {
			v := Increase2()
			fmt.Printf("goroutine %d: increase done: %d\n", id, v)
		}(i)
	}
}

func main() {
	// showChanCase5()
	showChanCase52()
	time.Sleep(time.Second * 5)
}
