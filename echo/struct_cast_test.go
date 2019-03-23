package echo

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestStructTest(t *testing.T) {
	var a *A = &A{
		a: 1,
		b: 2,
	}
	a.method()

	var b = (&B{a: 3,
		b: 4})

	c := (*A)(unsafe.Pointer(b))
	c.method()
}

type A struct {
	a int
	b int
}

func (a *A) method() {
	fmt.Println(a.a)
}

type B struct {
	a int
	b int
}
