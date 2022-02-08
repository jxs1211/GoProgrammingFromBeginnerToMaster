package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func show() {
	r := strings.NewReader("hello, gopher!")
	lr := io.LimitReader(r, 4)
	n, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d bytes copied\n", n)
}

func main() {
	show()
}

func Show() {
	r := strings.NewReader("hello, gopher!\n")
	lr := io.LimitReader(r, 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
