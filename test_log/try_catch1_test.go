package test_log

import (
	"fmt"
	"testing"
)

func TestNormalException(t *testing.T) {
	MyFunc4(nil)
}

/**
 * panic异常测试
 */
func MyFunc4(obj interface{}) {
	//if _, ok := obj.(TestType); !ok {
	//	return
	//}
	fmt.Println(obj.(TestType).data)
}

// 测试的数据的类型
type TestType struct {
	data string
}
