package main

import "fmt"

// go tool compile -S asm.go >> asm.s
//go:noinline
func add_a_and_b(a, b int) int {
	return a + b
}

func main() {
	c := add_a_and_b(2, 3)
	fmt.Println(c)
}
