package main

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {

	var si *SI
	if si == nil {
		fmt.Println("si is nil")
	}
}
