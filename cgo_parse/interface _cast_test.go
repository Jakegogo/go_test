package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type SI struct {
	f1 int
}

func TestInterfaceCast(t *testing.T) {

	i := CreateI()[0]
	fmt.Println(i)

	i1 := CreateI()[0].(interface{})
	fmt.Println(i1.(SI))

	i2 := (interface{})(CreateI()[0])
	fmt.Println(i2.(SI))

}

func CreateI() []interface{} {
	si := SI{
		f1: 2,
	}
	return []interface{}{si}
}

func TestInterfacePtrCast(t *testing.T) {

	//i := CreateIPtr()[0]
	//iI := unsafe.Pointer(i.(uintptr))
	//fmt.Println(iI.(SI))

	i1 := CreateIPtr()[0]
	var iI1 interface{} = (unsafe.Pointer(i1.(uintptr)))
	fmt.Println((iI1).(*SI))

}

func CreateIPtr() []interface{} {
	si := SI{
		f1: 2,
	}
	return []interface{}{reflect.ValueOf(&si).Pointer()}
}
