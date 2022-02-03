package main

import "fmt"

func main() {
	var s = "中国人"

	for i, v := range s {
		fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	}
	fmt.Println("")

	//byte sequence of s: 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
	var sl = []byte{0xe4, 0xb8, 0xad, 0xe5, 0x9b, 0xbd, 0xe4, 0xba, 0xba}
	fmt.Println(string(sl))
	for _, v := range sl {
		fmt.Printf("0x%x ", v)
	}
	fmt.Println("\n")

	sl[3] = 0xd0
	sl[4] = 0xd6
	sl[5] = 0xb9
	sl2 := []byte{0xd6, 0xb9}
	fmt.Printf("%v - %v - %x - %b - %c\n", []rune(string(sl2)), string(sl2), sl2, sl2, sl2)
	for i, v := range string(sl) {
		fmt.Printf("%d %x\n", i, v)
	}
}
