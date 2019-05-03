package monkey

import (
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
func replaceFunction(from, to uintptr) (original []byte) {
	jumpData := jmpToFunctionValue(to)
	f := rawMemoryAccess(from, len(jumpData))
	original = make([]byte, len(f))
	copy(original, f)

	data := Uint64(original[0:8])
	fmt.Println("bytes ptr:",data)

	fmt.Printf("origin data: %s\n", hex.EncodeToString(original))

	// copy origin function
	copyf :=  rawMemoryAccess(from, 600)
	copyOrigin := make([]byte, 600)
	copy(copyOrigin, copyf)

	copyToLocation(placehlder, copyf)

	copyToLocation(from, jumpData)
	return
}
