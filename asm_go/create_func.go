package main

import (
	"fmt"
	"reflect"
)

func CreateFunc(
	fType reflect.Type,
	f func(args []reflect.Value) (results []reflect.Value)) (reflect.Value, reflect.Type, error) {
	if fType.Kind() != reflect.Func {
		panic("invalid input")
	}

	var ins, outs *[]reflect.Type
	ins = new([]reflect.Type)
	outs = new([]reflect.Type)

	for i := 0; i < fType.NumIn(); i++ {
		*ins = append(*ins, fType.In(i))
	}

	for i := 0; i < fType.NumOut(); i++ {
		*outs = append(*outs, fType.Out(i))
	}
	var variadic bool
	variadic = fType.IsVariadic()

	funcImpl, funcType := AllocateStackFrame(*ins, *outs, variadic, f)
	return funcImpl, funcType, nil
}

func AllocateStackFrame(
	ins []reflect.Type,
	outs []reflect.Type,
	variadic bool,
	f func(args []reflect.Value) (results []reflect.Value)) (reflect.Value, reflect.Type) {
	var funcType reflect.Type
	funcType = reflect.FuncOf(ins, outs, variadic)
	return reflect.MakeFunc(funcType, f), funcType
}

func AnyGeneric(args []reflect.Value) (result []reflect.Value) {

	fmt.Printf("\n AnyGeneric called %d \n", args[0].Interface())

	var returnValue reflect.Value
	returnValue = reflect.ValueOf(0)

	return []reflect.Value{returnValue}
}
