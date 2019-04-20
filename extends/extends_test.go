package main

import (
	"fmt"
	"testing"
)

type A struct {
}

func (a *A) Foo() {
	fmt.Println("A.Foo()")
}

func (a *A) Bar() {
	a.Foo()
}

type B struct {
	A
}

func (b *B) Foo() {
	fmt.Println("B.Foo()")
}

func TestExtends(t *testing.T) {
	b := B{A: A{}}
	b.Bar()
}

func TestExtends1(t *testing.T) {
	a := A{}
	a.Bar()
}
