package main

// go tool compile -S asm.go >> asm.s
func add_a_and_b(a, b int) int {
	return a + b
}

func main() {
	add_a_and_b(2, 3)
}
