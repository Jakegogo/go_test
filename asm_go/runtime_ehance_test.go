package main

import (
	"fmt"
	forcexport "go_test/forcexport"
	"testing"
)

func TestFindFunc(t *testing.T) {
	ptr, _ := forcexport.FindFuncWithName("math.Sqrt")
	fmt.Printf("Found pointer 0x%x\nNormal function: %s", ptr, "math.Sqrt")
}
