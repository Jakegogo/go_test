package gls

import (
	"fmt"
	"github.com/modern-go/gls"
	"testing"
	"time"
)

func TestGls2(t *testing.T) {
	// 无法使用
	gls.Set("key1", "value1")
	fmt.Println(gls.Get("key1"))

	go func() {
		gls.Set("key1", "value2")
		fmt.Println(gls.Get("key1"))
	}()

	<-time.After(2 * time.Second)
}
