package main

import (
	"fmt"
	"strings"
)

func concat(a, b int) string {
	return fmt.Sprintf("%d %d", a, b)
}

func concatTwo(x, y string) string {
	return x + " " + y
}

func concatSlice(s []string) string {
	return strings.Join(s, " ")
}

func main() {
	println(concat(1, 2))
	println(concatTwo("hello", "gopher"))
	println(concatSlice([]string{"hello", "gopher", "!"}))
}
