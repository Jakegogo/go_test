package other_test

import (
	"fmt"
	"path/filepath"
	"strings"
)

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
)

func main() {
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)

	fmt.Println(0.0)

	p := strings.TrimPrefix("*Dcache.Proxy", "*")
	li := strings.LastIndex(p, ".")
	fmt.Printf("Dcache.Proxy"[:li] + "\n")

	fmt.Println(filepath.Dir("*test/Dcache.Proxy"))
}
