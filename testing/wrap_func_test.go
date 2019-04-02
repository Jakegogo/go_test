package main

import (
	"fmt"
	"reflect"
	"testing"
)

type (
	// Type of function being wrapped
	sumFuncT func(int, int) int

	// Type of the wrapper function
	wrappedSumFuncT func(sumFuncT, ...interface{}) int
)

// Wrapper of any function
// First element of array is the function being wrapped
// Other elements are arguments to the function
func genericWrapper(in []reflect.Value) []reflect.Value {
	// this is the place to do something useful in the wrapper
	arguments := in[1].Interface().([]interface{})
	reflectV := make([]reflect.Value, len(arguments))
	for i := 0; i < len(arguments); i++ {
		reflectV[i] = reflect.ValueOf(arguments[i])
	}
	return in[0].Call(reflectV)
}

// Creates wrapper function and sets it to the passed pointer to function
func createWrapperFunction(function interface{}) {
	fn := reflect.ValueOf(function).Elem()
	v := reflect.MakeFunc(reflect.TypeOf(function).Elem(), genericWrapper)
	fn.Set(v)
}

func TestWrapFunc(t *testing.T) {
	var wrappedSumFunc wrappedSumFuncT

	createWrapperFunction(&wrappedSumFunc)

	// The function being wrapped itself
	sumFunc := func(a int, b int) int {
		return a + b
	}

	result := wrappedSumFunc(sumFunc, 1, 3)
	fmt.Printf("Result is %v", result)
}
