package classtemplate

import (
	"fmt"
	"testing"
)

type Base struct {
}

func (b *Base) Method() {
	b.OnMethodCall()
}

func (b *Base) OnMethodCall() {
	fmt.Println("base method called")
}

type SubClass struct {
	Base
}

func (b *SubClass) OnMethodCall() {
	fmt.Println("sub method called")
}

func TestTemplate(t *testing.T) {
	var s = &SubClass{}
	s.Method()
}
