package main

import (
	"fmt"
	"reflect"
)

type Interface interface {
	M1()
	M2()
}

type T struct {
	Interface
}

func (T) M3() {}

func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemTyp := v.Elem()

	n := elemTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemTyp)
		return
	}

	fmt.Printf("%s's method set:\n", elemTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", elemTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

func main() {
	DumpMethodSet((*Interface)(nil))
	var t T
	var pt *T
	DumpMethodSet(&t)
	DumpMethodSet(&pt)
}
