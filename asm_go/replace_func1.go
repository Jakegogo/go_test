package main

import (
	"bou.ke/monkey"
	"fmt"
	forcexport "go_test/forcexport"
	"reflect"
)



func main() {
	//fmt.Print("ok")

	dynamicFunc1(fmt.Print, func(origin reflect.Value, args []reflect.Value) []reflect.Value {
		fmt.Println("called fmt.Print, args:",args)
		return origin.Call(args)
	})

	fmt.Print("ok")
}

func dynamicFunc1(fn interface{}, proxy func(origin reflect.Value, args []reflect.Value) []reflect.Value) {
	// monkey劫持函数调用
	//var originPrintLn func(a ...interface{}) (n int, err error)
	var originPrintLn = reflect.MakeFunc(reflect.TypeOf(fn), func(args []reflect.Value) (results []reflect.Value) {
		return []reflect.Value{}
	}).Interface()

	dynamicFunc := reflect.MakeFunc(reflect.TypeOf(fn), func(args []reflect.Value) (results []reflect.Value) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("time4 call error:", err)
				panic(err)
			}
		}()

		return proxy(reflect.ValueOf(originPrintLn), args)
	}).Interface()

	_ = dynamicFunc


	patchGuard := monkey.Patch(fn, dynamicFunc)
	// 构造原先方法
	fmt.Println("OrignUintptr is:", fmt.Sprintf("0x%x", patchGuard.OrignUintptr))
	forcexport.CreateFuncForCodePtrWithType(&originPrintLn, patchGuard.OrignUintptr, fn)

	fmt.Println("proxy ok")

}
