package main

import (
	"fmt"
	. "github.com/agiledragon/gomonkey"
	"go_test/forcexport"
	"math/rand"
)

func main() {

	var fastrand func() uint32
	forceexport.GetFunc(&fastrand, "runtime.fastrand")

	fmt.Println(fastrand())

	ApplyFunc(fastrand, func() uint32 {
		return 1991
	})

	fmt.Println(fastrand())
	fmt.Println(fastrand())

	myRecords := make(map[string][]string)
	myRecords["testKey"] = append(myRecords["testKey"], "a value")
	myRecords["testKey"] = append(myRecords["testKey"], "another value")
	myRecords["testKey"] = append(myRecords["testKey"], "just another value")
	myRecords["newKey"] = append(myRecords["newKey"], "and another value")
	myRecords["newKey"] = append(myRecords["newKey"], "and another")
	myRecords["lastKey"] = append(myRecords["lastKey"], "and another")
	myRecords["lastKey"] = append(myRecords["lastKey"], "and the last one")
	myRecords["lastKey"] = append(myRecords["lastKey"], "and the last one")
	for k, v := range myRecords {
		fmt.Printf("k: %s, v: %v\n", k, v)
	}

	fmt.Println("-------------")

	for k, v := range myRecords {
		fmt.Printf("k: %s, v: %v\n", k, v)
	}

	v := rand.Intn(100)
	fmt.Println(v)
	v = rand.Intn(100)
	fmt.Println(v)
	v = rand.Intn(100)
	fmt.Println(v)

}
