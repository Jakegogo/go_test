package other_test

import "fmt"

func main() {

	var a map[string]interface{}

	var p = &a
	var q = p

	a = make(map[string]interface{})
	(*q)["s"] = "k"
	fmt.Println(a)

}
