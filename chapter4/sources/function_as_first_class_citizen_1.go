package main

import "fmt"

type binaryCalcFunc func(int, int) int

func Show() {
	var i interface{} = binaryCalcFunc(func(x, y int) int { return x + y })
	c := make(chan func(int, int) int, 10)
	fns := []binaryCalcFunc{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x * y },
		func(x, y int) int { return x / y },
		func(x, y int) int { return x % y },
	}

	c <- func(x, y int) int {
		return x * y
	}

	fmt.Println(fns[0](5, 6))
	f := <-c
	fmt.Println(f(7, 10))
	v, ok := i.(binaryCalcFunc)
	if !ok {
		fmt.Println("type assertion error")
		return
	}

	fmt.Println(v(17, 7))
}

type AddFunc func(int, int) int

func Show2() {
	var i interface{} = AddFunc(func(x, y int) int { return x + y })
	c := make(chan func(int, int) int, 10)
	funcs := []AddFunc{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x * y },
		func(x, y int) int { return x / y },
	}
	c <- func(x, y int) int { return x * y }
	res := funcs[0](5, 6)
	fmt.Println(res)
	f := <-c
	fmt.Println(f(7, 10))
	v, ok := i.(AddFunc)
	if !ok {
		fmt.Println("type assert error")
		return
	}
	fmt.Println(v(17, 7))
}

func main() {
	Show()
	Show2()
}
