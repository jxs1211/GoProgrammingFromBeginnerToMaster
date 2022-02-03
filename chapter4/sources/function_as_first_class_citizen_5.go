// inspired by https://github.com/go-functional/core
package main

import (
	"fmt"
)

type InFactor interface {
	Map(fn func(int) int) InFactor
}

type Calculator struct {
	ints []int
}

func (c *Calculator) Map(fn func(int) int) InFactor {
	sl := make([]int, len(c.ints))
	for i, v := range c.ints {
		sl[i] = fn(v)
	}
	return &Calculator{ints: sl}
}

func NewCalculator(s []int) *Calculator {
	return &Calculator{ints: s}
}

func main() {

	s := []int{1, 2, 3}
	c := NewCalculator(s)

	plus := func(i int) int {
		return i + 10
	}

	mapped1 := c.Map(plus)
	fmt.Println("mapped1: ", mapped1)

	times := func(i int) int {
		return i * 3
	}

	mapped2 := mapped1.Map(times)
	fmt.Println("mapped2: ", mapped2)
	fmt.Println("composite map: ", c.Map(plus).Map(times))
	fmt.Println("original c: ", c)
}

func Show() {
	// 原切片
	intSlice := []int{1, 2, 3, 4}
	fmt.Printf("init a functor from int slice: %#v\n", intSlice)
	f := NewIntSliceFunctor(intSlice)
	fmt.Printf("original functor: %+v\n", f)

	mapperFunc1 := func(i int) int {
		return i + 10
	}

	mapped1 := f.Fmap(mapperFunc1)
	fmt.Printf("mapped functor1: %+v\n", mapped1)

	mapperFunc2 := func(i int) int {
		return i * 3
	}
	mapped2 := mapped1.Fmap(mapperFunc2)
	fmt.Printf("mapped functor2: %+v\n", mapped2)
	fmt.Printf("original functor: %+v\n", f) // 原functor没有改变
	fmt.Printf("composite functor: %+v\n", f.Fmap(mapperFunc1).Fmap(mapperFunc2))
}

type IntSliceFunctor interface {
	Fmap(fn func(int) int) IntSliceFunctor
}

type intSliceFunctorImpl struct {
	ints []int
}

func (isf intSliceFunctorImpl) Fmap(fn func(int) int) IntSliceFunctor {
	newInts := make([]int, len(isf.ints))
	for i, elt := range isf.ints {
		retInt := fn(elt)
		newInts[i] = retInt
	}
	return intSliceFunctorImpl{ints: newInts}
}

func NewIntSliceFunctor(slice []int) IntSliceFunctor {
	return intSliceFunctorImpl{ints: slice}
}
