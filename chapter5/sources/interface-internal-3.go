package main

import (
	"fmt"
	"unsafe"
)

type T int

func (t T) Error() string {
	return "bad error"
}

type T2 int
type In interface {
	String() string
}

func (t T2) String() string {
	return fmt.Sprintf("%d", t)
}

// go run interface-internal-3.go dumpinterface.go
func main() {
	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)

	dumpEface(eif)
	dumpItabOfIface(unsafe.Pointer(&err))
	dumpDataOfIface(err)

	var err2 In = T2(5)
	dumpDataOfIface(err2)

	// var eif3 interface{} = errors.New("5")
	// var err3 error = errors.New("5")

	// dumpEface(eif3)
	// dumpItabOfIface(unsafe.Pointer(&err3))
	// dumpDataOfIface(err)

}
