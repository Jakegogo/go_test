package main

import (
	"fmt"
	"reflect"
)

type A1 struct {
}

func (a *A1) Mehtod() {

}

func main() {
	a1 := &A1{}
	fmt.Println(reflect.ValueOf(a1.Mehtod).Addr())
}
