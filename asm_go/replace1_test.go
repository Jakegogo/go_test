package main

import (
	"testing"
	"unsafe"
)

func a1() int { return 1 }
func b1() int { return 2 }

func rawMemoryAccess1(b uintptr) []byte {
	return (*(*[0xFF]byte)(unsafe.Pointer(b)))[:]
}

func assembleJump(f func() int) []byte {
	funcVal := *(*uintptr)(unsafe.Pointer(&f))
	return []byte{
		0x48, 0xC7, 0xC2,
		byte(funcVal >> 0),
		byte(funcVal >> 8),
		byte(funcVal >> 16),
		byte(funcVal >> 24), // MOV rdx, funcVal
		0xFF, 0x22,          // JMP [rdx]
	}
}

func replace(orig, replacement func() int) {
	bytes := assembleJump(replacement)
	functionLocation := **(**uintptr)(unsafe.Pointer(&orig))
	window := rawMemoryAccess1(functionLocation)

	copy(window, bytes)
}

func TestReplace1(t *testing.T) {
	replace(a1, b1)
	print(a1())
}
