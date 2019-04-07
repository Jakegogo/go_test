package main

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func(int) int

func TestCreateFunc(t *testing.T) {
	var funcVar fn
	var funcType reflect.Type
	funcType = reflect.TypeOf(funcVar)
	v, _, err := CreateFunc(funcType, AnyGeneric)
	if err != nil {
		fmt.Printf("Error creating function %v\n", err)
	}

	input := 42

	ins := []reflect.Value([]reflect.Value{reflect.ValueOf(input)})
	outs := v.Call(ins)
	for i := range outs {
		fmt.Printf("%+v\n", outs[i].Interface())
	}
}
