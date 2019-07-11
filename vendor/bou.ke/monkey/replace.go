package monkey

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"golang.org/x/arch/x86/x86asm"
	"log"
	"reflect"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

var placeHolderOffset uintptr

// from is a pointer to the actual function
// to is a pointer to a go funcvalue
func replaceFunction(name string, from, to, orign, placehlder uintptr) (originalBytes *[]byte, orignFunc unsafe.Pointer, originPtr uintptr) {

	fmt.Println("from is:", fmt.Sprintf("0x%x", from))
	fmt.Println("placehlder is:", fmt.Sprintf("0x%x", placehlder))
	redirect1 := *(*int32)(unsafe.Pointer(from))
	fmt.Printf("redirect1 is:0x%x\n", redirect1)

	redirect := *(*uintptr)(unsafe.Pointer(from))
	fmt.Printf("before from redirect to 0x%x %d\n", redirect, redirect)

	jumpData := jmpToFunctionValue(to)

	ins, err := x86asm.Decode(jumpData, 64)
	if err == nil {
		fmt.Print("jump ins:", ins)
		ins, err := x86asm.Decode(jumpData[ins.Len:], 64)
		if err == nil {
			fmt.Print(" ", ins.String(), "\n")
		}
	}

	f := rawMemoryAccess(from, len(jumpData))
	original := make([]byte, len(f))
	copy(original, f)

	fmt.Printf("origin data: %s\n", hex.EncodeToString(original))

	// copy origin function
	funcSize, err := GetFuncSize(64, from, false)
	if err != nil {
		log.Println("GetFuncSize error", err)
		funcSize = 1024
	}
	fmt.Println("func size is", funcSize)

	copyf := rawMemoryAccess(from, funcSize)
	copyOrigin := make([]byte, funcSize)
	copy(copyOrigin, copyf)


	//placehlder = atomic.LoadUintptr(&placeHolderOffset)

	// replace replative address to placehlder
	replaceRelativeAddr(from, copyOrigin, placehlder, funcSize)

	copyToLocation(placehlder, copyOrigin)

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


func replaceRelativeAddr(from uintptr, copyOrigin []byte, placehlder uintptr, funcSize int) {

	coped := make([]byte, len(copyOrigin))
	copy(coped, copyOrigin)

	startAddr := (uint64)(from)
	for pos := 0; pos < len(copyOrigin); {

		endPos := pos + 16
		if endPos > len(copyOrigin) {
			endPos = len(copyOrigin)
		}

		code := copyOrigin[pos:endPos]

		ins, err := x86asm.Decode(code, 64)

		if err == nil && ins.Opcode != 0 {
			fmt.Printf("0x%x:\t%s\t\t%s\n", startAddr+(uint64)(pos), ins.Op, ins.String())

			replaceAll(ins, pos, copyOrigin, funcSize, startAddr, placehlder)

		}
		pos = pos + ins.Len
	}

	printCode(copyOrigin)

	fmt.Println("============================")

	for pos := 0; pos < len(coped); {

		endPos := pos + 16
		if endPos > len(coped) {
			endPos = len(coped)
		}

		code := coped[pos:endPos]

		ins, err := x86asm.Decode(code, 64)

		if err == nil && ins.Opcode != 0 {
			fmt.Printf("0x%x:\t%s\t\t%s\n", startAddr+(uint64)(pos), ins.Op, ins.String())


			replaceSeveral(ins, pos, startAddr, placehlder, coped, from)

		}
		pos = pos + ins.Len
	}

	printCode(coped)
}

func printCode(copyOrigin []byte) {
	fmt.Println("============================")
	fmt.Println("After:")
	for pos := 0; pos < len(copyOrigin); {

		endPos := pos + 16
		if endPos > len(copyOrigin) {
			endPos = len(copyOrigin)
		}

		code := copyOrigin[pos:endPos]

		ins, _ := x86asm.Decode(code, 64)
		fmt.Println(ins.String())
		pos = pos + ins.Len
	}
}

func replaceSeveral(ins x86asm.Inst, pos int, startAddr uint64, placehlder uintptr, copyOrigin []byte, from uintptr) {
	if "CALL" == ins.Op.String() {

		addrArgs := ins.Args[0].String()
		if strings.HasPrefix(addrArgs, ".+") || strings.HasPrefix(addrArgs, ".-") {
			relativeAddr, err := strconv.ParseUint(addrArgs[2:], 0, 32)
			if strings.HasPrefix(addrArgs, ".-") {
				relativeAddr = - relativeAddr
			}
			if err == nil {

				// TODO 递归处理,如果是绝对地址

				//copy(copyf[offset:offset+4], relativeAddr)
				offset := pos + 1
				fmt.Println((int)(startAddr - (uint64)(placehlder) + relativeAddr))
				binary.LittleEndian.PutUint32(copyOrigin[offset:offset+4], (uint32)(startAddr-(uint64)(placehlder)+relativeAddr))

				ins, err := x86asm.Decode(copyOrigin[pos:pos+ins.Len], 64)
				if err == nil {
					fmt.Printf("after replace: 0x%x:\t%s\t\t%s\n", startAddr+(uint64)(pos), ins.Op, ins.String())
				}
			} else {
				panic(err)
			}
		}

	}
	if "JMP" == ins.Op.String() {
		addrArgs := ins.Args[0].String()
		fmt.Println(addrArgs)
		if strings.HasPrefix(addrArgs, ".-") {
			jmpAddr, err := strconv.ParseUint(addrArgs[2:], 0, 32)
			if err == nil {
				if (uint64)(startAddr+(uint64)(pos)+(uint64)(ins.Len)-jmpAddr) == (uint64)(from) {
					//offset := pos + 1
					//binary.LittleEndian.PutUint32(copyOrigin[offset:offset+4], (uint32)(placehlder))
				}
				// TODO JMP外部函数处理
			}
		}
	}
	// TODO MOVQ
	if "MOV" == ins.Op.String() {
		addrArgs := ins.Args[0].String()
		if strings.Contains(addrArgs, "RIP+") || strings.Contains(addrArgs, "RIP-") {
			offset := pos + 3
			relativeAddr := binary.LittleEndian.Uint32(copyOrigin[offset : offset+4])
			if strings.Contains(addrArgs, "RIP-") {
				relativeAddr = - relativeAddr
			}

			fmt.Println((int)(startAddr - (uint64)(placehlder) + (uint64)(relativeAddr)))
			binary.LittleEndian.PutUint32(copyOrigin[offset:offset+4], (uint32)(startAddr-(uint64)(placehlder)+(uint64)(relativeAddr)))

			ins, err := x86asm.Decode(copyOrigin[pos:pos+ins.Len], 64)
			if err == nil {
				fmt.Printf("after replace: 0x%x:\t%s\t\t%s\n", startAddr+(uint64)(pos), ins.Op, ins.String())
			}
		}

		addrArgs1 := ins.Args[1].String()
		if strings.Contains(addrArgs1, "RIP+") || strings.Contains(addrArgs1, "RIP-") {
			offset := pos + 3
			relativeAddr := binary.LittleEndian.Uint32(copyOrigin[offset : offset+4])
			if strings.Contains(addrArgs1, "RIP-") {
				relativeAddr = - relativeAddr
			}

			fmt.Println((int)(startAddr - (uint64)(placehlder) + (uint64)(relativeAddr)))
			binary.LittleEndian.PutUint32(copyOrigin[offset:offset+4], (uint32)(startAddr-(uint64)(placehlder)+(uint64)(relativeAddr)))

			ins, err := x86asm.Decode(copyOrigin[pos:pos+ins.Len], 64)
			if err == nil {
				fmt.Printf("after replace: 0x%x:\t%s\t\t%s\n", startAddr+(uint64)(pos), ins.Op, ins.String())
			}
		}
	}
	// TODO LEAQ
	if "LEA" == ins.Op.String() {
		addrArgs := ins.Args[1].String()
		if strings.Contains(addrArgs, "RIP+") || strings.Contains(addrArgs, "RIP-") {
			offset := pos + 3
			relativeAddr := binary.LittleEndian.Uint32(copyOrigin[offset : offset+4])
			if strings.Contains(addrArgs, "RIP-") {
				relativeAddr = - relativeAddr
			}
			fmt.Println((int)(startAddr - (uint64)(placehlder) + (uint64)(relativeAddr)))
			binary.LittleEndian.PutUint32(copyOrigin[offset:offset+4], (uint32)(startAddr-(uint64)(placehlder)+(uint64)(relativeAddr)))

			ins, err := x86asm.Decode(copyOrigin[pos:pos+ins.Len], 64)
			if err == nil {
				fmt.Printf("after replace: 0x%x:\t%s\t\t%s\n", startAddr+(uint64)(pos), ins.Op, ins.String())
			}
		}
		fmt.Print(addrArgs)
	}
}

func replaceAll(ins x86asm.Inst, pos int, copyOrigin []byte, funcSize int, startAddr uint64, placehlder uintptr) {
	// 需要替换偏移地址
	if ins.PCRelOff > 0 {

		offset := pos + ins.PCRelOff

		var isAdd = true
		for i := 0; i < len(ins.Args); i++ {
			arg := ins.Args[i]
			if arg == nil {
				break
			}
			addrArgs := arg.String()
			if strings.HasPrefix(addrArgs, ".-") || strings.Contains(addrArgs, "RIP-") {
				isAdd = false
			}
		}

		relativeAddr := decodeAddress(copyOrigin[offset:offset+ins.PCRel], ins.PCRel)
		if !isAdd && relativeAddr > 0 {
			relativeAddr = - relativeAddr
		}


		fmt.Println("added", (int)(relativeAddr)+pos+ins.Len)
		if (isAdd && (int)(relativeAddr)+pos+ins.Len >= funcSize) ||
			(!isAdd && (int)(relativeAddr)+pos+ins.Len < 0) {

			if ins.Op.String() == "CALL" {
				fmt.Println((int64)(startAddr)-(int64)(placehlder), startAddr, placehlder, int32(relativeAddr))
			}

			encodeAddress(copyOrigin[offset:offset+ins.PCRel], ins.PCRel, relativeAddr, (int)(startAddr)-(int)(placehlder))

			ins, err := x86asm.Decode(copyOrigin[pos:pos+ins.Len], 64)
			if err == nil {
				fmt.Printf("after replace: \t%s\t\t%s\n", ins.Op, ins.String())
			}
		} else {
			if ins.Op.String() == "CALL" {
				fmt.Println((int)(relativeAddr)+pos+ins.Len, funcSize, (int)(relativeAddr)+pos+ins.Len)
				fmt.Println("called")
			}
		}

	}
}


func encodeAddress(bytes []byte, len int, val int, add int) {
	if len == 1 {
		bytes[0] = byte(int(val) + add)
	} else if len == 2 {
		LittleEndian.PutInt16(bytes, int16(val) + int16(add))
	} else if len == 4 {
		LittleEndian.PutInt32(bytes, int32(val) + int32(add))
	} else if len == 8 {
		LittleEndian.PutInt64(bytes, int64(val) + int64(add))
	}
}

func decodeAddress(bytes []byte, len int) int {
	if len == 1 {
		return int(bytes[0])
	} else if len == 2 {
		return int(LittleEndian.Int16(bytes))
	} else if len == 4 {
		fmt.Println(LittleEndian.Int32(bytes))
		return int(LittleEndian.Int32(bytes))
	} else if len == 8 {
		return int(LittleEndian.Int64(bytes))
	}
	return 0
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
