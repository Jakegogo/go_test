// Automatically generated by AspectGo (github.com/Jakegogo/aspectgo)
// DO NOT EDIT MANUALLY

package aspect

import "fmt"
import "go_test/another"

var global int = 3

func AsyncCall() {
	go Call(1)

	// nothing

	// go func
	var a = 1
	//var c = make(chan map[string]bool, 5)
	//go func(c chan map[string]bool,param1 int) {
	//	fmt.Println("async ", a)
	//}(c, 1)
	fmt.Println("sync", a, global, GlobalVarIns.key, another.GlobalVarIns1)
}

func Call(param1 int) {
	fmt.Println("async ")
}