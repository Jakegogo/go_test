package test_nil

import (
	"fmt"
	"testing"
)

func TestNilCast(t *testing.T) {
	m := make(map[string]interface{})
	fmt.Println(m["ok"].(string))
}


func TestNilFunc(t *testing.T) {
	a := []func() error{nil}
	a[0]()
}