package main

import (
	"bou.ke/monkey"
	"fmt"
	forcexport "go_test/forcexport"
)

func time5(a int) int {
	result := a * 5
	fmt.Print("called time2\n")
	return result
}

func time6(a int) int {
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
	return result
}

func main() {
	patchGuard := monkey.Patch1(time5, func(a int) int {
		fmt.Printf("\npatch called %d\n", a)
		return 0
	}, time6)

	var originTime2 func(a int) int
	forcexport.CreateFuncForCodePtr(&originTime2, patchGuard.OrignUintptr-1)
	fmt.Println("result is:", originTime2(3))
	fmt.Println("actual result is:", time5(3))
}
