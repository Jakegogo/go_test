package main

import (
	"fmt"
	"unsafe"
)

func a1(a int) int {
	result := a * 4
	fmt.Print("called a1\n")
	return result
}

func main1() {
	f := a1
	f(1)
	a1(1)
	fmt.Printf("0x%x\n", **(**uintptr)(unsafe.Pointer(&f)))
}
