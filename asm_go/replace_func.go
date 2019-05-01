package main

import (
	"bou.ke/monkey"
	"fmt"
	forcexport "go_test/forcexport"
)

func time5(a int) int {
	result := a * 11
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
	var originTime2 func(a int) int
	patchGuard := monkey.Patch1(time5, func(a int) int {
		fmt.Println("orign result is:", originTime2(3))
		return a * 6
	}, time6)

	forcexport.CreateFuncForCodePtr(&originTime2, patchGuard.OrignUintptr)
	fmt.Println("wraper result is:", time5(3))
	fmt.Println("placeholder result is:", time6(3))
}
