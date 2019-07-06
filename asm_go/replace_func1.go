package main

import (
	"bou.ke/monkey"
	"fmt"
	forcexport "go_test/forcexport"
	"net"
	"reflect"
	"syscall"
)



func main() {
	dynamicFunc1()
}

var printLn func(a ...interface{}) (n int, err error)

func dynamicFunc1() {
	// monkey劫持函数调用

	var originPrintLn func(a ...interface{}) (n int, err error)

	dynamicFunc := reflect.MakeFunc(reflect.TypeOf(fmt.Print), func(args []reflect.Value) (results []reflect.Value) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("time4 call error:", err)
				panic(err)
			}
		}()
		r, e := originPrintLn("args:", args[0].Interface())
		fmt.Println(r, e)
		return []reflect.Value{reflect.ValueOf(r), reflect.ValueOf(e)}
	}).Interface()

	_ = dynamicFunc


	patchGuard := monkey.Patch1(fmt.Print, dynamicFunc, getPlaceHolder4)
	// 构造原先方法
	fmt.Println("OrignUintptr is:", fmt.Sprintf("0x%x", patchGuard.OrignUintptr))
	forcexport.CreateFuncForCodePtrWithType(&originPrintLn, patchGuard.OrignUintptr, fmt.Print)

	//time4(3)
	//time.Sleep(time.Second * 200)
	//originTime4(3)

	fmt.Print("ok")

	fmt.Println("ok")

}



// Network file descriptor.
type netFD struct {
	pfd FD

	// immutable until Close
	family      int
	sotype      int
	isConnected bool // handshake completed or use of association with peer
	net         string
	laddr       net.Addr
	raddr       net.Addr
}

// FD is a file descriptor. The net and os packages use this type as a
// field of a larger type representing a network connection or OS file.
type FD struct {
	// Lock sysfd and serialize access to Read and Write methods.
	fdmu fdMutex

	// System file descriptor. Immutable until Close.
	Sysfd int

	// I/O poller.
	pd pollDesc

	// Writev cache.
	iovecs *[]syscall.Iovec

	// Semaphore signaled when file is closed.
	csema uint32

	// Non-zero if this file has been set to blocking mode.
	isBlocking uint32

	// Whether this is a streaming descriptor, as opposed to a
	// packet-based descriptor like a UDP socket. Immutable.
	IsStream bool

	// Whether a zero byte read indicates EOF. This is false for a
	// message based socket connection.
	ZeroReadIsEOF bool

	// Whether this is a file rather than a network socket.
	isFile bool
}

type fdMutex struct {
	state uint64
	rsema uint32
	wsema uint32
}

type pollDesc struct {
	runtimeCtx uintptr
}


func getPlaceHolder4(a int) int {
	result := a * 4
	result = result * 5
	result = result * 6
	result = result * 7
	result = result * 8
	result = result * 9
	result = result * 10
	result = result * 11
	result = result * 12
	result = result * 13
	result = result * 14
	result = result * 15
	result = result * 16
	result = result * 17
	result = result * 18
	result = result * 19
	result = result * 20
	result = result * 21
	result = result * 22
	result = result * 23
	result = result * 24
	result = result * 25
	result = result * 26

	result = result * 27
	result = result * 28
	result = result * 29
	result = result * 30
	result = result * 31
	result = result * 32
	result = result * 33
	result = result * 34
	result = result * 35
	result = result * 36
	result = result * 37
	result = result * 38
	result = result * 39
	result = result * 40
	result = result * 41
	result = result * 42
	result = result * 43
	result = result * 44
	result = result * 45
	result = result * 46
	result = result * 47
	result = result * 48

	result = result * 49
	result = result * 50
	result = result * 51
	result = result * 52
	result = result * 53
	result = result * 54
	result = result * 55
	result = result * 56
	result = result * 57
	result = result * 58
	result = result * 59
	result = result * 60
	result = result * 61
	result = result * 62

	result = result * 27
	result = result * 28
	result = result * 29
	result = result * 30
	result = result * 31
	result = result * 32
	result = result * 33
	result = result * 34
	result = result * 35
	result = result * 36
	result = result * 37
	result = result * 38
	result = result * 39

	result = result * 50
	result = result * 51
	result = result * 52
	result = result * 53
	result = result * 54
	result = result * 55
	result = result * 56
	result = result * 57
	result = result * 58
	result = result * 59
	result = result * 60
	result = result * 61
	result = result * 62

	result = result * 56
	result = result * 57
	result = result * 58
	result = result * 59
	result = result * 60
	result = result * 61

	fmt.Print("called getPlaceHolder3\n")
	fmt.Print("called getPlaceHolder \n")
	return result
}