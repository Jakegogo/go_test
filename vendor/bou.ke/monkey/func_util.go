package monkey

import (
	"bytes"
	"golang.org/x/arch/x86/x86asm"
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


func GetFuncSize(mode int, start uintptr, minimal bool) (int, error) {

	prologueLen := len(funcPrologue)
	code := rawMemoryAccess(start, 16) // instruction takes at most 16 bytes

	/* prologue is not required
	if !bytes.Equal(funcPrologue, code[:prologueLen]) { // not valid function start or invalid prologue
		return 0, errors.New(fmt.Sprintf("no func prologue, addr:0x%x", start))
	}
	*/

	int3_found := false
	curLen := 0

	for {
		inst, err := x86asm.Decode(code, mode)
		if err != nil || (inst.Opcode == 0 && inst.Len == 1 && inst.Prefix[0] == x86asm.Prefix(code[0])) {
			return curLen, nil
		}

		if inst.Len == 1 && code[0] == 0xcc {
			// 0xcc -> int3, trap to debugger, padding to function end
			if minimal {
				return curLen, nil
			}
			int3_found = true
		} else if int3_found {
			return curLen, nil
		}

		curLen = curLen + inst.Len
		code = rawMemoryAccess(start+uintptr(curLen), 16) // instruction takes at most 16 bytes

		if bytes.Equal(funcPrologue, code[:prologueLen]) {
			return curLen, nil
		}
	}

	return 0, nil
}
