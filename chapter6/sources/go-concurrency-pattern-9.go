package main

import (
	"fmt"
	"sync"
	"time"
)

func newNumGenerator(start, count int) <-chan int {
	c := make(chan int)
	go func() {
		for i := start; i < start+count; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func filterOdd(in int) (int, bool) {
	time.Sleep(time.Microsecond * 1000)
	if in%2 != 0 {
		return 0, false
	}
	return in, true
}

func square(in int) (int, bool) {
	return in * in, true
}

func spawnGroup(name string, num int, f func(int) (int, bool), in <-chan int) <-chan int {
	groupOut := make(chan int)
	var outSlice []chan int
	for i := 0; i < num; i++ {
		out := make(chan int)
		go func(i int) {
			name := fmt.Sprintf("%s-%d:", name, i)
			fmt.Printf("%s begin to work...\n", name)

			for v := range in {
				r, ok := f(v)
				if ok {
					out <- r
				}
			}
			close(out)
			fmt.Printf("%s work done\n", name)
		}(i)
		outSlice = append(outSlice, out)
	}

	// Fan-in
	//
	// out --\
	//        \
	// out ---- --> groupOut
	//        /
	// out --/
	//
	go func() {
		var wg sync.WaitGroup
		for _, out := range outSlice {
			wg.Add(1)
			go func(out <-chan int) {
				defer wg.Done()
				for v := range out {
					groupOut <- v
				}

			}(out)
		}
		wg.Wait()
		close(groupOut)
	}()

	return groupOut
}

// pipline
//   1     newNumGenerator
// 1 1 1   filterOdd
//  1 1    square
//   1     out
//

func main() {
	start := time.Now()

	in := newNumGenerator(1, 1000)
	out := spawnGroup("square", 30, square, spawnGroup("filterOdd", 30, filterOdd, in))

	time.Sleep(3 * time.Second)

	for v := range out {
		fmt.Println(v)
	}
	println("elasped: ", time.Since(start)/1000000)
}
