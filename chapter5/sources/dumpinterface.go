package main

import (
	"fmt"
	"unsafe"
)

// type T int

// func (t T) Error() string {
// 	return "bad error"
// }

// type T2 int
// type In interface {
// 	String() string
// }

// func (t T2) String() string {
// 	return fmt.Sprintf("%d", t)
// }

const ptrSize = unsafe.Sizeof(uintptr(0))

type typeAlg struct {
	hash  func(unsafe.Pointer, uintptr) uintptr
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}

type tflag uint8
type nameOff int32
type typeOff int32

type _type struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32
	tflag      tflag
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        *typeAlg
	gcdata     *byte
	str        nameOff
	ptrToThis  typeOff
}

type imethod struct {
	name nameOff
	ityp typeOff
}

type name struct {
	bytes *byte
}

type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}

type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}

type eface struct {
	_type *_type
	data  unsafe.Pointer
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

// fit for go 1.13.x version
func dumpEface(i interface{}) {
	ptrToEface := (*eface)(unsafe.Pointer(&i))
	fmt.Printf("eface: %+v\n", *ptrToEface)

	if ptrToEface._type != nil {
		// dump _type info
		fmt.Printf("\t _type: %+v\n", *(ptrToEface._type))
	}

	if ptrToEface.data != nil {
		// dump data
		switch i.(type) {
		case int:
			println("i'type is int")
			dumpInt(ptrToEface.data)
		case float64:
			println("i'type is float64")
			dumpFloat64(ptrToEface.data)
		case T:
			println("i'type is T")
			dumpT(ptrToEface.data)

		// other cases ... ...
		default:
			fmt.Printf("\t data: unsupported type\n")
		}
	}
	fmt.Printf("\n")
}

func dumpInt(dataOfEface unsafe.Pointer) {
	var p *int = (*int)(dataOfEface)
	fmt.Printf("\t data: %d\n", *p)
}
func dumpFloat64(dataOfEface unsafe.Pointer) {
	var p *float64 = (*float64)(dataOfEface)
	fmt.Printf("\t data: %f\n", *p)
}

func dumpItabOfIface(ptrToIface unsafe.Pointer) {
	println("dumpItabOfIface")
	p := (*iface)(ptrToIface)
	fmt.Printf("iface: %+v\n", *p)

	if p.tab != nil {
		// dump itab
		fmt.Printf("\t itab: %+v\n", *(p.tab))
		// dump inter in itab
		fmt.Printf("\t\t inter: %+v\n", *(p.tab.inter))

		// dump _type in itab
		fmt.Printf("\t\t _type: %+v\n", *(p.tab._type))

		// dump fun in tab
		funPtr := unsafe.Pointer(&(p.tab.fun))
		fmt.Printf("\t\t fun: [")
		for i := 0; i < len((*(p.tab.inter)).mhdr); i++ {
			tp := (*uintptr)(unsafe.Pointer(uintptr(funPtr) + uintptr(i)*ptrSize))
			fmt.Printf("0x%x(%d),", *tp, *tp)
		}
		fmt.Printf("]\n")
	}
}

func dumpDataOfIface(i interface{}) {
	println("dumpDataOfIface")
	// this is a trick as the data part of eface and iface are same
	ptrToEface := (*iface)(unsafe.Pointer(&i))

	if ptrToEface.data != nil {
		// dump data
		switch i.(type) {
		case int:
			dumpInt(ptrToEface.data)
		case float64:
			dumpFloat64(ptrToEface.data)
		case T:
			dumpT(ptrToEface.data)
		case T2:
			dumpT2(ptrToEface.data)
		// other cases ... ...

		default:
			fmt.Printf("\t data: unsupported type\n")
		}
	}
	fmt.Printf("\n")
}

func dumpT2(p unsafe.Pointer) {
	var v = (*T2)(p)
	fmt.Printf("\t %+v\n", *v)
}

func dumpT(dataOfIface unsafe.Pointer) {
	var p *T = (*T)(dataOfIface)
	fmt.Printf("\t data: %+v\n", *p)
}

// func dumpError(dataOfIface unsafe.Pointer) {
// 	var p *error = (*error)(dataOfIface)
// 	fmt.Printf("\t data: %+v\n", *p)
// }

// func main() {
// 	var eif interface{} = T(5)
// 	var err error = T(5)
// 	println("eif:", eif)
// 	println("err:", err)
// 	println("eif = err:", eif == err)

// 	dumpEface(eif)
// 	dumpItabOfIface(unsafe.Pointer(&err))
// 	dumpDataOfIface(err)

// 	var err2 In = T2(5)
// 	dumpDataOfIface(err2)
// }
