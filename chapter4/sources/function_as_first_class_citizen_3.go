package main

import "fmt"

type Adder interface {
	Add(int, int) int
}

type AddFunc func(int, int) int

func (a AddFunc) Add(x, y int) int {
	return a(x, y)
}

func main() {
	var adder Adder = AddFunc(func(x, y int) int { return x + y })
	fmt.Println(adder.Add(5, 6))
}

type BinaryAdder interface {
	Add(int, int) int
}

type MyAdderFunc func(int, int) int

func (f MyAdderFunc) Add(x, y int) int {
	return f(x, y)
}

func MyAdd(x, y int) int {
	return x + y
}
