package main

import "time"

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	print("\tgoroutine-", id, ": idCheck ok\n")
	return idCheckTmCost
}

func bodyCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	print("\tgoroutine-", id, ": bodyCheck ok\n")
	return bodyCheckTmCost
}

func xRayCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	print("\tgoroutine-", id, ": xRayCheck ok\n")
	return xRayCheckTmCost
}

func airportSecurityCheck(id int) int {
	print("goroutine-", id, ": airportSecurityCheck ...\n")
	total := 0

	total += idCheck(id)
	total += bodyCheck(id)
	total += xRayCheck(id)

	print("goroutine-", id, ": airportSecurityCheck ok\n")
	return total
}

func start(id int, f func(int) int, queue <-chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		total := 0
		for {
			_, ok := <-queue
			if !ok {
				c <- total
				return
			}
			total += f(id)
		}
	}()
	return c
}

func max(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return n
}

func main() {
	total := 0
	passengers := 300
	c := make(chan struct{})
	c1 := start(1, airportSecurityCheck, c)
	c2 := start(2, airportSecurityCheck, c)
	c3 := start(3, airportSecurityCheck, c)

	// goroutines := 4
	// res := make([]<-chan int, goroutines)
	// for i := 0; i < goroutines; i++ {
	// 	cn := start(i+1, airportSecurityCheck, c)
	// 	res = append(res, cn)
	// }

	for i := 0; i < passengers; i++ {
		c <- struct{}{}
	}
	close(c)
	// res := make([]int, 3)
	// res = append(res, <-c1)
	// res = append(res, <-c2)
	// res = append(res, <-c3)
	// total = max(res...)

	// ints := make([]int, len(res))
	// for i, v := range res {
	// 	// if len(v) <= 0 {
	// 	// 	continue
	// 	// }
	// 	func(i int, v <-chan int) {
	// 		ints[i] = <-v
	// 	}(i, v)

	// }
	// total = max(ints...)

	total = max(<-c1, <-c2, <-c3)
	println("total time cost:", total)
}
