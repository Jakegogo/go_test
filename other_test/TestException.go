package main

import (
	"fmt"
	"log"
	"os"
)

var (
	logger *log.Logger
)

func init() {
	f, _ := os.Create("log")
	logger = log.New(f, "", log.LstdFlags)
}

func main() {
	fmt.Println("Begin.")

	errorFunc()

	fmt.Println("End.") // NO End.
}

func errorFunc() {
	bytes := make([]byte, 0)

	var ResBody *[]byte // 请求体
	ResBody = &bytes
	fmt.Println(*ResBody)

	var s Struct1
	fmt.Println(s.Msg())

	var i *int = nil
	//*i = 3
	fmt.Println(*i)

	var sl = []string{"a", "b"}
	//defer recoverPanic()
	fmt.Println(sl[3])
}

func recoverPanic() {
	if rec := recover(); rec != nil {
		err := rec.(error)
		logger.Printf("Unhandled error: %v\n", err.Error())
		fmt.Fprintf(os.Stderr, "Program quit unexpectedly; please check your logs\n", err.Error())
		//os.Exit(1)
	}
}

type Struct1 struct {
}

func (s Struct1) Msg() string {
	return "ok"
}
