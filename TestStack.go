package main

import (
	"go_test/xerrs"
)

import "fmt"

func main() {

	//....

	if data, err := MyFunc(); err != nil {
		xerrs.SetData(err, "some_key", "VALUE") // set custom error value

		fmt.Println(data)
		//....

		// Details returns cause error, mask if specified, and the stack (accepting the maximum stack height as parameter)
		// In this example only 5 last stack function calls will be printed out
		xerrs.Print(err)
		fmt.Println(xerrs.GetData(err, "some_key"))

		//....
	}

}

//func MyFunc() (interface{}, error) {
//	return "result", xerrs.New("error occur");
//}

func MyFunc() (interface{}, error) {
	return "result", xerrs.Errorf("wraped error message", xerrs.New("error occur"))
}
