package main

import "testing"

func TestInterface(t *testing.T) {

	si := &SI{}
	Call(si)
}

func Call(u interface{}) {
	u.(S).Ok()
}

type S interface {
	Ok()
}

type SI struct {
}

func (s SI) Ok() {
	panic("implement me")
}
