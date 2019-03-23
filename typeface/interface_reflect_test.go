package main

import (
	"reflect"
	"testing"
)

func TestInterfaceReflect(t *testing.T) {
	s := &TestService{}
	callMethod(s)
}

func callMethod(s interface{}) {
	method := reflect.ValueOf(s).MethodByName("Add")
	method.Call([]reflect.Value{reflect.ValueOf(0), reflect.ValueOf(1)})
}
