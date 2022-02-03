package main

import "fmt"

func foo(b ...byte) {
	fmt.Println(string(b))
}

func main() {
	b := []byte{}
	b = append(b, "hello"...)
	fmt.Printf("%b\n%o\n%d\n0x%x\n%s\n", b, b, b, b, b)

	foo(b...)
	// foo("hello")
}
