package array_test

import "testing"

var a [10000]int
var sl = a[:]

func arrayLoop() {

	var sum int
	for _, v := range a {
		sum += v
	}
}

func arrayPointLoop() {
	var sum int
	for _, v := range &a {
		sum += v
	}
}

func arraySliceLoop() {
	var sum int
	for _, v := range sl {
		sum += v
	}
}

func BenchmarkArrayLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayLoop()
	}
}

func BenchmarkArrayPointLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayPointLoop()
	}
}

func BenchmarkArraySliceLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arraySliceLoop()
	}
}
