package forceexport

import (
	"fmt"
	"go_test/xerrs"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

// GetFunc gets the function defined by the given fully-qualified name. The
// outFuncPtr parameter should be a pointer to a function with the appropriate
// type (e.g. the address of a local variable), and is set to a new function
// value that calls the specified function. If the specified function does not
// exist, outFuncPtr is not set and an error is returned.
func GetFunc(outFuncPtr interface{}, name string) error {
	codePtr, err := FindFuncWithName(name)
	if err != nil {
		return err
	}
	CreateFuncForCodePtr(outFuncPtr, codePtr)
	return nil
}

// Convenience struct for modifying the underlying code pointer of a function
// value. The actual struct has other values, but always starts with a code
// pointer.
type Func struct {
	codePtr uintptr
}

// CreateFuncForCodePtr is given a code pointer and creates a function value
// that uses that pointer. The outFun argument should be a pointer to a function
// of the proper type (e.g. the address of a local variable), and will be set to
// the result function value.
func CreateFuncForCodePtr(outFuncPtr interface{}, codePtr uintptr) *Func {
	outFuncVal := reflect.ValueOf(outFuncPtr).Elem()
	// Use reflect.MakeFunc to create a well-formed function value that's
	// guaranteed to be of the right type and guaranteed to be on the heap
	// (so that we can modify it). We give a nil delegate function because
	// it will never actually be called.
	newFuncVal := reflect.MakeFunc(outFuncVal.Type(), nil)
	// Use reflection on the reflect.Value (yep!) to grab the underling
	// function value pointer. Trying to call newFuncVal.Pointer() wouldn't
	// work because it gives the code pointer rather than the function value
	// pointer. The function value is a struct that starts with its code
	// pointer, so we can swap out the code pointer with our desired value.
	funcValuePtr := reflect.ValueOf(newFuncVal).FieldByName("ptr").Pointer()
	funcPtr := (*Func)(unsafe.Pointer(funcValuePtr))
	funcPtr.codePtr = codePtr
	outFuncVal.Set(newFuncVal)
	return funcPtr
}

// FindFuncWithName searches through the moduledata table created by the linker
// and returns the function's code pointer. If the function was not found, it
// returns an error. Since the data structures here are not exported, we copy
// them below (and they need to stay in sync or else things will fail
// catastrophically).
func FindFuncWithName(name string) (uintptr, error) {
	for moduleData := &Firstmoduledata; moduleData != nil; moduleData = moduleData.next {
		for _, ftab := range moduleData.ftab {
			f := (*runtime.Func)(unsafe.Pointer(&moduleData.pclntable[ftab.funcoff]))
			if f == nil {
				continue
			}
			var fName string
			err := xerrs.Try(func() {
				fName = f.Name()
			})
			if err != nil {
				fmt.Println(xerrs.Details(err, 50))
			}
			if strings.HasSuffix(fName, name) {
				fmt.Printf("FindFuncWithName found %s\n", name)
				return f.Entry(), nil
			}
		}
	}
	fmt.Printf("FindFuncWithName not found %s", name)
	return 0, fmt.Errorf("Invalid function name: %s\n", name)
}

// Everything below is taken from the runtime package, and must stay in sync
// with it.

//go:linkname Firstmoduledata runtime.firstmoduledata
var Firstmoduledata Moduledata

type Moduledata struct {
	pclntable    []byte
	ftab         []Functab
	filetab      []uint32
	findfunctab  uintptr
	minpc, maxpc uintptr

	text, etext           uintptr
	noptrdata, enoptrdata uintptr
	data, edata           uintptr
	bss, ebss             uintptr
	noptrbss, enoptrbss   uintptr
	end, gcdata, gcbss    uintptr

	// Original type was []*_type
	typelinks []interface{}

	modulename string
	// Original type was []modulehash
	modulehashes []interface{}

	gcdatamask, gcbssmask Bitvector

	next *Moduledata
}

type Functab struct {
	entry   uintptr
	funcoff uintptr
}

type Bitvector struct {
	n        int32 // # of bits
	bytedata *uint8
}