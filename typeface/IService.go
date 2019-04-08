package main

/*
DO NOT EDIT!
This code was generated automatically using github.com/hexdigest/typeface
The original type "TestService" can be found in go_test/typeface package
You can generate mock for this interface using github.com/gojuno/minimock:

minimock -i go_test/typeface.IService -o ./
*/

//IService contains exportable methods signatures of the go_test/typeface.TestService
type IService interface {
	Add(a int, b int) (int, error, int)
}
