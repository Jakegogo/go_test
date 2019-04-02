package main

import (
	"bou.ke/monkey"
	"fmt"
	reflect "reflect"
	"runtime"
	"testing"
	"time"
)

/*
将创建Func封装， 非reflect.Func类型会panic
当然makeFunc的闭包函数表达式类型是固定的，可以查阅一下文档。
细读文档的reflect.Value.Call()方法。
*/
func MakeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)
	if rf.Kind() != reflect.Func {
		panic("非Reflect.Func")
	}
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func time2(a int) int {
	result := a * 2
	return result
}

func TestMakeFunc(t *testing.T) {
	timedToo := MakeTimedFunction(time2).(func(int) int)
	time2Val := timedToo(5)
	fmt.Println(time2Val)
}

func TestMakeFuncWithPatch(t *testing.T) {

	time2(6)

	timedToo := MakeTimedFunction(time2).(func(int) int)
	monkey.Patch(time2, timedToo)
	time2Val := timedToo(5)
	fmt.Println(time2Val)
}
