package gls

import (
	"fmt"
	"github.com/jtolds/gls"
	"testing"
	"time"
)

var (
	mgr = gls.NewContextManager()
)

func TestGls(t *testing.T) {
	go mgr.SetValues(gls.Values{"test_key": "test_val"}, func() {
		fmt.Println(mgr.GetValue("test_key"))
	})

	go mgr.SetValues(gls.Values{"test_key1": "test_val1"}, func() {
		fmt.Println(mgr.GetValue("test_key"))
		fmt.Println(mgr.GetValue("test_key1"))
	})

	<-time.After(2 * time.Second)
}

func TestBenchmark(t *testing.T) {
	mgr := gls.NewContextManager()
	for i := 0; i < 1000000; i++ {
		mgr.SetValues(gls.Values{"test_key": "test_val"}, func() {
			mgr.SetValues(gls.Values{"test_key2": "test_val2"}, func() {})
		})
	}
}
