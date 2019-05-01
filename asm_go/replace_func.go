package main

import (
	"bou.ke/monkey"
	"fmt"
	forcexport "go_test/forcexport"
	"reflect"
)

func time4(a int) int {
	result := a * 11
	return result
}

func time5(a int) int {
	result := a * 22
	return result
}

func main() {
	var originTime4 func(a int) int
	var originTime5 func(a int) int
	// monkey劫持函数调用
	patchGuard := monkey.Patch1(time4, func(a int) int {
		// 调用原先的方法
		fmt.Println("orign result1 is:", originTime4(3))
		return a * 6
	}, getPlaceHolder(1))
	// 构造原先方法
	forcexport.CreateFuncForCodePtr(&originTime4, patchGuard.OrignUintptr)
	fmt.Println("orign result1 is:", originTime4(3))

	patchGuard1 := monkey.Patch1(time5, func(a int) int {
		// 调用原先的方法
		fmt.Println("orign result2 is:", originTime5(3))
		return a * 6
	}, getPlaceHolder(2))
	// 构造原先方法
	forcexport.CreateFuncForCodePtr(&originTime5, patchGuard1.OrignUintptr)

	// 调用包装后的方法
	fmt.Println("wraper result1 is:", time4(3))
	fmt.Println("wraper result2 is:", time5(3))
}

func getPlaceHolder(a int) interface{} {
	var placeholder func(a int64) (int64, string)
	return reflect.MakeFunc(reflect.TypeOf(placeholder), func(args []reflect.Value) (results []reflect.Value) {
		result := a * 4
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5
		result = result * 5

		fmt.Print("called time4\n")
		return []reflect.Value{reflect.ValueOf(result)}
	}).Interface()
}
