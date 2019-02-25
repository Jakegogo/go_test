package echo

import (
	"fmt"
	"testing"
)

func TestStructTest(t *testing.T) {
	var a A
	a.method()

	//var b = (&B{})
	//a = b.(A)
}

type A struct {
	a int
	b int
}

func (a *A) method() {
	fmt.Println("a")
}

type B struct {
	A
}

func (a *B) method() {
	fmt.Println("b")
}
