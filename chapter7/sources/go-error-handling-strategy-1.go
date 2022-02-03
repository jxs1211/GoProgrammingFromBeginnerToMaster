package main

import (
	"errors"
	"fmt"
)

var ErrSentinel = errors.New("the underlying sentinel error")

type errString struct {
	s string
}

func (e *errString) Error() string {
	return e.s
}

func MyNew(s string) *errString {
	return &errString{s}
}

func MyNew2(s string) errString {
	return errString{s}
}

var ErrExist = errors.New("file already exists")

var ErrExist2 = MyNew("file already exists")

func ShowCompareStrutPointsAndValues() {
	error1 := MyNew("file already exists")
	error2 := MyNew("file already exists")
	if error1 != error2 {
		fmt.Printf("struct pointer compare: %v != %v\n", error1, error2)
	} else {
		fmt.Printf("struct pointer compare: %v == %v\n", error1, error2)
	}

	err1 := MyNew2("file already exists")
	err2 := MyNew2("file already exists")
	if err1 != err2 {
		fmt.Printf("struct value compare: %v != %v", err1, err2)
	} else {
		fmt.Printf("struct value compare: %v == %v", err1, err2)
	}

}

func main() {
	ShowCompareStrutPointsAndValues()
	err1 := fmt.Errorf("wrap err1: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err2: %w", err1)
	if errors.Is(err2, ErrSentinel) {
		println("err is ErrSentinel")
		return
	}

	println("err is not ErrSentinel")
}
