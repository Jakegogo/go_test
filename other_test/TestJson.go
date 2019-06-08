package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString := `{
	"foo": {
		"b": "1",
		"a": "2",
		"c": "3"
	}
}`
	o := make(map[string]map[string]string)
	err := json.Unmarshal([]byte(jsonString), &o)
	if err != nil {
		panic(err)
	}

	fmt.Println("range:")
	for _, m := range o {
		for k, v := range m {
			fmt.Printf("%s, %s\n", k, v)
		}
	}

	fmt.Println("range:")
	for _, m := range o {
		for k, v := range m {
			fmt.Printf("%s, %s\n", k, v)
		}
	}

	bytes, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	fmt.Println("")
	fmt.Println("json:")
	fmt.Printf("%s\n", bytes)
}
