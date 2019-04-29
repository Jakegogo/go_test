package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func a() int { return 1 }

func TestPointCut1() {

	f := a

	ptr1 := *(*uintptr)(unsafe.Pointer(&f))
	fmt.Printf("0x%x\n", ptr1)

	ptr2 := **(**uintptr)(unsafe.Pointer(&f))
	fmt.Printf("0x%x\n", ptr2)

	var f1 interface{}
	f1 = a
	fmt.Printf("0x%x\n", reflect.ValueOf(f1).Pointer())

	f2 := getPtr(reflect.ValueOf(f))
	fmt.Printf("0x%x\n", f2)

}

func getPtr(v reflect.Value) unsafe.Pointer {
	return (*value1)(unsafe.Pointer(&v)).ptr
}

type value1 struct {
	_   uintptr
	ptr unsafe.Pointer
}
