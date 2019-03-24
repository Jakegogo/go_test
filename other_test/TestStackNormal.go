package other_test

import (
	"go_test/xerrs"
)

import "fmt"

func main() {

	//....

	if _, err := MyFunc5(); err != nil {
		err := xerrs.Extend(err) // extend error 无法还原出原始普通异常的堆栈,建议使用xerrs.New

		fmt.Println(xerrs.Cause(err))

		//....

		// Details returns cause error, mask if specified, and the stack (accepting the maximum stack height as parameter)
		// In this example only 5 last stack function calls will be printed out
		xerrs.Print(err)

		//....
	}

}

//func MyFunc() (interfacetest{}, error) {
//	return "result", xerrs.New("error occur");
//}

func MyFunc5() (interface{}, error) {
	return "result", fmt.Errorf("error occur")
}
