package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2, 3))
}

func Sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
