package main

import (
	"bou.ke/monkey"
	"fmt"
	forcexport "go_test/forcexport"
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
	fmt.Println("before orign result1 is:", originTime4(3))

	patchGuard1 := monkey.Patch1(time5, func(a int) int {
		// 调用原先的方法
		fmt.Println("orign result2 is:", originTime5(3))
		return a * 6
	}, getPlaceHolder2(2))
	// 构造原先方法
	forcexport.CreateFuncForCodePtr(&originTime5, patchGuard1.OrignUintptr)

	// 调用包装后的方法
	fmt.Println("wraper result1 is:", time4(3))
	fmt.Println("wraper result2 is:", time5(3))
}

func getPlaceHolder(a int) interface{} {
	return func(a int64) (int64, string) {
		result := a * 4
		result = result * 5
		result = result * 6
		result = result * 7
		result = result * 8
		result = result * 9
		result = result * 10
		result = result * 11
		result = result * 12
		result = result * 13
		result = result * 14
		result = result * 15
		result = result * 16
		result = result * 17
		result = result * 18
		result = result * 19
		result = result * 20
		result = result * 21
		result = result * 22
		result = result * 23
		result = result * 24
		result = result * 25
		result = result * 26

		result = result * 27
		result = result * 28
		result = result * 29
		result = result * 30
		result = result * 31
		result = result * 32
		result = result * 33
		result = result * 34
		result = result * 35
		result = result * 36
		result = result * 37
		result = result * 38
		result = result * 39
		result = result * 40
		result = result * 41
		result = result * 42
		result = result * 43
		result = result * 44
		result = result * 45
		result = result * 46
		result = result * 47
		result = result * 48

		result = result * 49
		result = result * 50
		result = result * 51
		result = result * 52
		result = result * 53
		result = result * 54
		result = result * 55
		result = result * 56
		result = result * 57
		result = result * 58
		result = result * 59
		result = result * 60
		result = result * 61
		result = result * 62
		fmt.Print("called getPlaceHolder\n")
		return result, ""
	}
}

func getPlaceHolder2(a int) interface{} {
	return func(a int64) (int64, string) {
		result := a * 4
		result = result * 5
		result = result * 6
		result = result * 7
		result = result * 8
		result = result * 9
		result = result * 10
		result = result * 11
		result = result * 12
		result = result * 13
		result = result * 14
		result = result * 15
		result = result * 16
		result = result * 17
		result = result * 18
		result = result * 19
		result = result * 20
		result = result * 21
		result = result * 22
		result = result * 23
		result = result * 24
		result = result * 25
		result = result * 26

		result = result * 27
		result = result * 28
		result = result * 29
		result = result * 30
		result = result * 31
		result = result * 32
		result = result * 33
		result = result * 34
		result = result * 35
		result = result * 36
		result = result * 37
		result = result * 38
		result = result * 39
		result = result * 40
		result = result * 41
		result = result * 42
		result = result * 43
		result = result * 44
		result = result * 45
		result = result * 46
		result = result * 47
		result = result * 48

		result = result * 49
		result = result * 50
		result = result * 51
		result = result * 52
		result = result * 53
		result = result * 54
		result = result * 55
		result = result * 56
		result = result * 57
		result = result * 58
		result = result * 59
		result = result * 60
		result = result * 61
		result = result * 62
		fmt.Print("called getPlaceHolder2\n")
		return result, ""
	}
}
