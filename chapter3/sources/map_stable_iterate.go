package main

import "fmt"

func doIteration(sl []int, m map[int]int) {
	fmt.Printf("{ ")
	for _, k := range sl {
		v, ok := m[k] // keep the order
		if !ok {
			continue
		}
		fmt.Printf("[%d, %d] ", k, v)
	}
	fmt.Printf("}\n")
}

func main() {
	var sl []int
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}

	for k, _ := range m {
		sl = append(sl, k)
	}
	fmt.Println(sl)
	for i := 0; i < 3; i++ {
		doIteration(sl, m)
	}
}
