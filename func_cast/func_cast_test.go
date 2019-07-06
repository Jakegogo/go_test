package func_cast

import (
	"fmt"
	"go_test/forcexport"
	"reflect"
	"testing"
	"unsafe"
)

type Hanlder func(Ctx) error



type Ctx interface {
	Ok()
}

type Ctx1 interface {
	Ok()
}


func TestFuncCast(t *testing.T) {
	var h Hanlder

	var h1 = func(Ctx1) error {
		fmt.Println("h1 called.")
		return nil
	}

	var h1I interface{} = &h1

	fmt.Println((uintptr(unsafe.Pointer(&h1))))
	fmt.Println(reflect.ValueOf(&h1).Pointer())
	fmt.Println(reflect.ValueOf(h1I).Pointer())


	h = *(*Hanlder)(unsafe.Pointer(reflect.ValueOf(h1I).Pointer()))

	h(nil)


}


func TestFuncCast1(t *testing.T) {

	var h1 Hanlder = func(Ctx) error {
		fmt.Println("h1 called.")
		return nil
	}

	var h1I interface{} = h1

	fmt.Println((uintptr(unsafe.Pointer(&h1))))
	fmt.Println(reflect.ValueOf(&h1).Pointer())
	fmt.Println(reflect.ValueOf(h1I).Pointer())

	hPtr := reflect.ValueOf(h1I).Pointer()
	var H2 func(Ctx1) error
	forceexport.CreateFuncForCodePtr(&H2, hPtr)


	H2(nil)


}


func TestFuncCast2(t *testing.T) {
	var h1 Hanlder = func(Ctx) error {
		fmt.Println("h1 called.")
		return nil
	}

	doProxyCall(h1)
}

func doProxyCall(h1 interface{}) {
	func1 := reflect.MakeFunc(reflect.TypeOf(h1), func(args []reflect.Value) (results []reflect.Value) {
		fmt.Println(args[0].Interface())
		return reflect.ValueOf(h1).Call(args)
	}).Interface()
	func1.(Hanlder)(nil)
}