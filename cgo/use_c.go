package main

// void SayHello(const char* s1);
import "C"
import "fmt"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
	fmt.Println("ok")
}
