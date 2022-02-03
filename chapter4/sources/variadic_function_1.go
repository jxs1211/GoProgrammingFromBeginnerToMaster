package main

import "fmt"

func sums(ints ...int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func dump2(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func Show() {
	s := []int{1, 2, 3}
	fmt.Println(sums(s...))
	fmt.Println(sums(4, 5, 6))
}

func Show2() {
	s := []interface{}{1, 2, 3}
	dump2(s...)
}

func main() {
	// Show()
	Show2()
}
