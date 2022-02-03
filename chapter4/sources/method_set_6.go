package main

type Interface interface {
	M1()
	M2()
}

type Interface2 interface {
	M1()
	M2()
}

type T struct {
	Interface
	Interface2
}

func (T) M1() {
	println("T's M1")
}

func (T) M2() {
	println("T's M2")
}

type S struct{}

func (S) M1() {
	println("S's M1")
}
func (S) M2() {
	println("S's M2")
}

type S2 struct{}

func (S2) M1() {
	println("S's M1")
}
func (S2) M2() {
	println("S's M2")
}

func main() {
	var t = T{
		Interface: S{},
	}

	t.M1()
	t.M2()
}
