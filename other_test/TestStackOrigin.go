package other_test

import (
	"errors"
	"go_test/xerrs"
)

import "fmt"

func main() {

	//....

	defer func() {
		if err := recover(); err != nil {
			err := xerrs.Catch(err.(error)) // catch error

			// Details returns cause error, mask if specified, and the stack (accepting the maximum stack height as parameter)
			// In this example only 5 last stack function calls will be printed out
			fmt.Println(xerrs.Details(err, 15))
			fmt.Println(xerrs.GetData(err, "some_key"))
		}
	}()

	MyFunc3()

}

func MyFunc3() (interface{}, error) {
	panic(errors.New("an error occured."))
}
