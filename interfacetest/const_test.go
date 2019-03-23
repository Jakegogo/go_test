package main

import (
	"fmt"
	"testing"
)

func TestConst(t *testing.T) {
	a := newConst()

	fmt.Println(a)
	fmt.Println(*a)

	*a = 100

	fmt.Println(a)
	fmt.Println(*a)

	b := newConst()
	fmt.Println(b)
	fmt.Println(*b)

}

func newConst() *int {
	var defaultStatusCode = 200 // 每次会分配新的内存地址
	return &defaultStatusCode
}
