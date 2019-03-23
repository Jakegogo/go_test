package main

import (
	"fmt"
	"go_test/typeassert"
	"reflect"
	"strconv"
	"testing"
)

type Test struct {
}

func (t *Test) SayHello(name string, d int, s Test, s1 typeassert.S1) bool {
	return true
}

func FuncAnalyse(m interface{}) {
	//Reflection type of the underlying data of the interface
	x := reflect.TypeOf(m)

	numIn := x.NumIn()   //Count inbound parameters
	numOut := x.NumOut() //Count outbounding parameters

	fmt.Println("Method:", x.String())
	fmt.Println("Variadic:", x.IsVariadic()) // Used (<type> ...) ?
	fmt.Println("Package:", x.PkgPath())

	for i := 0; i < numIn; i++ {
		inV := x.In(i)
		in_Kind := inV.Kind() //func
		fmt.Printf("\nParameter IN: "+strconv.Itoa(i)+"\nKind: %v\nName: %v\n", in_Kind, inV.Name())
		if in_Kind == reflect.Struct {
			numField := inV.NumField()
			for j := 0; j < numField; j++ {
				field := inV.Field(j)
				fmt.Printf("\t%s - %v\n", field.Name, field.Type)
			}
		}
		fmt.Printf("-----------")
	}
	for o := 0; o < numOut; o++ {
		returnV := x.Out(0)
		return_Kind := returnV.Kind()
		fmt.Printf("\nParameter OUT: "+strconv.Itoa(o)+"\nKind: %v\nName: %v\n", return_Kind, returnV.Name())
	}
}

func TestMethod_Reflect(t *testing.T) {
	var n *Test = &Test{}
	methodValue := n.SayHello
	FuncAnalyse(methodValue)
}
