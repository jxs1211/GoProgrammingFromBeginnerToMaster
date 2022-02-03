package main

import (
	"fmt"
	"io"
	"reflect"
)

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
	DumpMethodSet((*io.Writer)(nil))
	DumpMethodSet((*io.Reader)(nil))
	DumpMethodSet((*io.Closer)(nil))
	DumpMethodSet((*io.ReadWriter)(nil))
	DumpMethodSet((*io.ReadWriteCloser)(nil))
}
