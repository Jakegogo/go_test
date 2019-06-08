package main

import "fmt"

type Elm struct {
}

func (Elm) Nothing() {
	panic("implement me")
}

type IElm interface {
	Nothing()
}

func main() {

	if getElm() == nil {
		fmt.Println("getElm() is nil")
	}

	a, _ := getElm().([]IElm)
	if getElm() == nil {
		a = nil
	}

	fmt.Println(a)
	fmt.Println(getElm())

	err, _ := getElm().(error)
	fmt.Println(err)
}

func getElm() interface{} {
	return nil
}
