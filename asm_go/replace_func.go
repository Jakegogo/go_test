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

func placeholder(a int) int {
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
	var originTime5 func(a int) int
	patchGuard := monkey.Patch1(time5, func(a int) int {
		fmt.Println("orign result is:", originTime5(3))
		return a * 6
	}, placeholder)

	forcexport.CreateFuncForCodePtr(&originTime5, patchGuard.OrignUintptr)
	fmt.Println("wraper result is:", time5(3))
	fmt.Println("placeholder result is:", placeholder(3))
}
