package other_test

import (
	"fmt"
	"go_test/xerrs"
	"runtime/debug"
)

func main() {

	fmt.Println("Begin.")
	// try..catch..异常捕获
	err := xerrs.Try(func() {
		MyFunc2()
	})

	if err != nil {
		xerrs.Print(err)
	}

	fmt.Println("End.")

	fmt.Println(string(debug.Stack()))
}

/**
 * panic异常测试
 */
func MyFunc2() {
	var sl = []string{"a", "b"}
	fmt.Println(sl[3])
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
