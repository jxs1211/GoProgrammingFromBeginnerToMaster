package main

import "fmt"

func Times(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func main() {

	timesTwo := Times(2)
	timesThree := Times(3)
	timesFour := Times(4)

	fmt.Println(timesTwo(5))
	fmt.Println(timesThree(5))
	fmt.Println(timesFour(5))
}

func Show() {
	timesTwo := partialTimes(2)
	timesThree := partialTimes(3)
	timesFour := partialTimes(4)
	fmt.Println(timesTwo(5))
	fmt.Println(timesThree(5))
	fmt.Println(timesFour(5))
}

func times(x, y int) int {
	return x * y
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}
