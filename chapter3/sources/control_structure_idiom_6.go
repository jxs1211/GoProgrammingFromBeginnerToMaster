package main

import (
	"fmt"
	"time"
)

func ShowBreakAtWhere() {
	exit := make(chan interface{})

	go func() {
		// loop:
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			case <-exit:
				fmt.Println("exiting...")
				// break loop
				break // break outside the select, not the for
			}
		}
		fmt.Println("exit!")
	}()

	time.Sleep(3 * time.Second)
	exit <- struct{}{}

	// wait child goroutine exit
	time.Sleep(3 * time.Second)
}

func ShowLabelLoop() {
	final_sl := [][]int{}
outLoop:
	for i := 0; i < 10; i++ {
		sl := []int{}
		fmt.Printf("sl len: %d sl cap: %d\n", len(sl), cap(sl))
		for j := 0; j < 10; j++ {
			if i >= 5 {
				break outLoop
			}
			if j >= 5 {
				continue outLoop
			}
			// fmt.Printf("%d * %d = %d\n", i, j, i*j)
			sl = append(sl, i*j)
		}
		fmt.Println(sl)
		final_sl = append(final_sl, sl)
	}
	fmt.Println(final_sl)
}

func main() {
	// ShowLabelLoop()
	sl := []int{}
	sl = append(sl, 1)
	fmt.Println(sl)
}
