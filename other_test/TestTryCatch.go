package main

import (
	"fmt"
	"go_test/xerrs"
	"runtime/debug"
)

func main() {

	tester := Tester{}

	fmt.Println("Begin.")
	// try..catch..异常捕获
	err := xerrs.Try(func() {
		tester.MyFunc2()
	})

	if err != nil {
		xerrs.Print(err)
	}

	fmt.Println("End.")

	fmt.Println(string(debug.Stack()))
}

type Tester struct {
}

/**
 * panic异常测试
 */
func (t *Tester) MyFunc2() {
	innerFun()
}

/**
 * panic异常测试
 */
func innerFun() {
	var c *Channel
	c.Attributes()

	//var a error = nil
	//fmt.Println(a)
	//
	//b:=a.(error)
	//fmt.Println(b)
	//
	//var sl = []string{"a", "b"}
	//fmt.Println(sl[3])
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

type Channel struct {
	hash string
}

func (m *Channel) Attributes() {
	r := "x" + m.hash
	fmt.Println(r)
}
