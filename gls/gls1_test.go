package gls

import (
	"fmt"
	"github.com/tylerb/gls"
	"testing"
	"time"
)

func TestGls1(t *testing.T) {
	gls.Set("key1", "value1")
	fmt.Println(gls.Get("key1"))

	go func() {
		gls.Set("key1", "value2")
		fmt.Println(gls.Get("key1"))
	}()

	<-time.After(2 * time.Second)
}

func TestGls1Performance(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		gls.Set("key1", "value1")
		gls.Get("key1")
	}
}
