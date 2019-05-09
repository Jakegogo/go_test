package monkey

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
)

func rawMemoryAccess(p uintptr, length int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
}

func pageStart(ptr uintptr) uintptr {
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}

// from is a pointer to the actual function
// to is a pointer to a go funcvalue
func replaceFunction(from, to, orign, placehlder uintptr) (originalBytes *[]byte, orignFunc unsafe.Pointer, originPtr uintptr) {

	fmt.Println("from is:", from)
	fmt.Println("placehlder is:", placehlder)
	redirect1 := *(*int32)(unsafe.Pointer(from))
	fmt.Println("redirect1 is:", redirect1)

	redirect := *(*uintptr)(unsafe.Pointer(from))
	fmt.Printf("before from redirect to 0x%x %d\n", redirect, redirect)

	jumpData := jmpToFunctionValue(to)

	f := rawMemoryAccess(from, len(jumpData))
	original := make([]byte, len(f))
	copy(original, f)

	data := Uint64(original[0:8])
	fmt.Println("bytes ptr:", data)

	fmt.Printf("origin data: %s\n", hex.EncodeToString(original))

	// copy origin function
	copyf := rawMemoryAccess(from, 470)
	copyOrigin := make([]byte, 470)
	copy(copyOrigin, copyf)

	copyToLocation(placehlder, copyf)

	copyToLocation(from, jumpData)
	fmt.Printf("from:%d, to:%d, orign:%d, originPtr:%d\n", from, to, orign, (uintptr)(unsafe.Pointer(&original[0])))

	mprotectCrossPage((uintptr)(unsafe.Pointer(&copyOrigin[0])), len(f), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)

	originJumpData := jmpToFunctionValue((uintptr)(reflect.ValueOf(original).Pointer()))
	fmt.Println(originJumpData)

	redirect = *(*uintptr)(unsafe.Pointer(from))
	fmt.Printf("after from redirect to 0x%x %d\n", redirect, redirect)

	redirect = *(*uintptr)(unsafe.Pointer(to))
	fmt.Printf("to redirect to 0x%x %d\n", redirect, redirect)

	return &copyOrigin, unsafe.Pointer(&original), placehlder
}

func Uint64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

// from is a pointer to the actual function
// to is a pointer to a go funcvalue
func replaceFunctionWithJump(from, orign uintptr, jumpData *[]byte) (original []byte, orignFunc unsafe.Pointer, originPtr uintptr) {
	f := rawMemoryAccess(from, len(*jumpData))
	original = make([]byte, len(f))
	copy(original, f)

	copyToLocation(from, *jumpData)
	fmt.Printf("from:%d, orign:%d, originPtr:%d\n", from, orign, (uintptr)(unsafe.Pointer(&original[0])))

	mprotectCrossPage((uintptr)(unsafe.Pointer(&original[0])), len(f), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)

	return original, unsafe.Pointer(&original), (uintptr)(unsafe.Pointer(&original[0]))
}
