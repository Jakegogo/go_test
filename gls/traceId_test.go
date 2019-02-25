package gls

import (
	"fmt"
	"testing"
	"time"
)

func TestTraceId(t *testing.T) {
	fmt.Println(GetTraceID())
	Clear()
	fmt.Println(GetTraceID())
}

func TestRuntinePass(t *testing.T) {
	fmt.Println(GetTraceID())
	context := GetContext()

	go func() {
		SetContext(context)
		fmt.Println(GetTraceID())
	}()

	time.Sleep(1 * time.Second)
}
