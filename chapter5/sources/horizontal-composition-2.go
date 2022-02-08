package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type UpperReader struct {
	r io.Reader
}

func NewUpperReader(r io.Reader) *UpperReader {
	return &UpperReader{r: r}
}

func (r *UpperReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}
	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, nil
}

func show() {
	r := NewUpperReader(io.LimitReader(strings.NewReader("hello, gopher!"), 5))
	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("\n %d bytes copied", n)
	fmt.Printf("\n %d bytes copied", n)

}

func main() {
	show()
}

func CapReader(r io.Reader) io.Reader {
	return &capitalizedReader{r: r}
}

type capitalizedReader struct {
	r io.Reader
}

func (r *capitalizedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}

	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, err
}

func Show() {
	r := strings.NewReader("hello, gopher!\n")
	r1 := CapReader(io.LimitReader(r, 4))
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		log.Fatal(err)
	}
}
