package main

import (
	"fmt"
	"net"
)

var C_s int = 9

func time4_b(a int) int {
	result := a * 11
	if result < 50 {
		return 3
	}
	if result < 100 {
		return 6
	}
	if result < 200 {
		return 9
	}
	C_s=3
	fmt.Println("yes")
	fmt.Println(C_s)
	//t6Result := 0
	//t6Result := time6fun.Call([]reflect.Value{reflect.ValueOf(0)})[0].Interface().(int)
	return result * time6_b(3) + 1
}

func trace_b(i int) {
	fmt.Println(i)
}

func time5_b(a int) int {
	result := a * 22
	fmt.Println("origin time5 called")
	return result
}

func time6_b(a int) int {
	result := a * 3
	fmt.Println("time6 called")
	return result + time7_b(2)
}

func time7_b(b int) int {
	result := b * 101
	net.Listen("tcp", "0.0.0.0:8090")
	fmt.Println("time7 called")
	return result + 1
}
