package main

import (
	"fmt"
	forcexport "go_test/forcexport"
	"math"
)

func main() {
	ptr, _ := forcexport.FindFuncWithName("math.Sqrt")
	fmt.Printf("Found pointer 0x%x\nNormal function: %s", ptr, math.Sqrt)
}
