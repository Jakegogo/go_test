package main

import "fmt"

type A string

func (a A) Error() string {
	return fmt.Sprintf("%s is an error", a)
}

func main() {
	a := A("hello")
	fmt.Printf("error is %s", a)
}
