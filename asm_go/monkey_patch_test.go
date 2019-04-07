package main

import (
	"bou.ke/monkey"
	"fmt"
	"reflect"
	"syscall"
	"testing"
	"unsafe"
)

type time2def func(a int) int

func time2(a int) int {
	result := a * 2
	fmt.Print("called time2")
	return result
}

func time3(a int) int {
	result := a * 3
	fmt.Print("called time3")
	return result
}

func TestMakeFuncWithCopyFunc(t *testing.T) {
	patchGuard := monkey.Patch(time2, func(a int) int {
		fmt.Printf("\npatch called %d\n", a)
		return 0
	})
	if patchGuard != nil {

	}

	//tfunc := time3
	//targetuinptr := **(**uintptr)(unsafe.Pointer(&tfunc))
	//fmt.Printf("targetuinptr is 0x%x\n", targetuinptr)

	monkey.PatchWithJump(time3, patchGuard.OrignBytes)

	time2(2)
	time3(2)

}

type Interface struct {
	typ     unsafe.Pointer
	funcPtr *Func
}

type Func struct {
	codePtr uintptr
}

func TestMakeFuncWithMakeFunc2(t *testing.T) {
	patchGuard := monkey.Patch(time2, func(a int) int {
		fmt.Printf("\npatch called %d\n", a)
		return 0
	})
	if patchGuard != nil {

	}

	time2(1)

	//func0 := (*runtime.Func)(unsafe.Pointer((patchGuard.OrignUintptr)))
	//fmt.Printf("func0 name:%s", func0.Name())

	funDef := func(a int) int { return 111 }
	var timeNowInterface interface{} = funDef
	timeNowInterfacePtr := (*Interface)(unsafe.Pointer(&timeNowInterface))

	timeNowInterfacePtr.funcPtr.codePtr = patchGuard.OrignUintptr
	sec := funDef(0)
	fmt.Println(sec)

	//funPointer := (unsafe.Pointer)(patchGuard.OrignUintptr)
	//funCast := (*func(a int) int)(funPointer)

	func1 := reflect.MakeFunc(reflect.TypeOf(time2), func(args []reflect.Value) (results []reflect.Value) {
		return []reflect.Value{reflect.ValueOf(0)}
	}).Interface()

	value1 := (*value)(unsafe.Pointer(&func1))

	pointer := value1.ptr

	funImpl := (*makeFuncImpl)(pointer)

	//(uintptr)(binary.LittleEndian.Uint32(patchGuard.OrignBytes))
	funImpl.code = (uintptr)(patchGuard.OrignUintptr)

	func1.(func(a int) int)(103)
}

func mprotectCrossPage(addr uintptr, length int, prot int) {
	pageSize := syscall.Getpagesize()
	for p := pageStart(addr); p < addr+uintptr(length); p += uintptr(pageSize) {
		page := rawMemoryAccess(p, pageSize)
		err := syscall.Mprotect(page, prot)
		if err != nil {
			panic(err)
		}
	}
}

func pageStart(ptr uintptr) uintptr {
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}

func rawMemoryAccess(p uintptr, length int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
}

func TestMakeFuncWithMakeFunc(t *testing.T) {
	patchGuard := monkey.Patch(time2, func(a int) int {
		fmt.Printf("\npatch called %d\n", a)
		return 0
	})
	if patchGuard != nil {

	}

	var funcVar time2def
	var funcType reflect.Type
	funcType = reflect.TypeOf(funcVar)
	v, _, err := CreateFunc(funcType, AnyGeneric)
	if err != nil {
		panic(err)
	}

	func1 := v.Interface()

	value1 := (*value)(unsafe.Pointer(&func1))

	pointer := value1.ptr

	funImpl := (*makeFuncImpl)(pointer)

	//(uintptr)(binary.LittleEndian.Uint32(patchGuard.OrignBytes))
	funImpl.code = (uintptr)(patchGuard.OrignUintptr)

	func1.(func(a int) int)(102)
}

func TestMakeFuncWithReplace(t *testing.T) {
	patchGuard := monkey.Patch(time2, func(a int) int {
		fmt.Printf("\npatch called %d\n", a)
		return 0
	})
	if patchGuard != nil {

	}

	//func1 := reflect.MakeFunc(reflect.TypeOf(time2), func(args []reflect.Value) (results []reflect.Value) {
	//
	//	return []reflect.Value{reflect.ValueOf(0)}
	//}).Interface()

	var funcVar time2def
	var funcType reflect.Type
	funcType = reflect.TypeOf(funcVar)
	v, _, err := CreateFunc(funcType, AnyGeneric)
	if err != nil {
		panic(err)
	}

	func1 := v.Interface()

	monkey.PatchWithJump(func1, patchGuard.OrignBytes)

	func1.(func(a int) int)(102)
}

func TestMakeFuncWithPatch(t *testing.T) {
	patchGuard := monkey.Patch(time2, func(a int) int {
		fmt.Printf("\npatch called %d\n", a)
		return 0
	})
	if patchGuard != nil {

	}

	// 直接调用orign字节指向的方法
	//var orignFunc func(a int) int
	//orignFunc = *(*func(a int) int)(patchGuard.OrignPtr)
	//reflect.ValueOf(orignFunc).Call([]reflect.Value{reflect.ValueOf(0)})

	//(orignFunc)(1)

	// 修改makefunc指针的方式
	func1 := reflect.MakeFunc(reflect.TypeOf(time2), func(args []reflect.Value) (results []reflect.Value) {

		return []reflect.Value{reflect.ValueOf(0)}
	})

	//func1.Interface().(func(a int) int)(3)

	value1 := (*value)(unsafe.Pointer(&func1))

	pointer := value1.ptr
	//p1 := (*int)(pointer)
	//*p1 = 21312

	funImpl := (*makeFuncImpl)(pointer)
	//orignPtr := *(*int)(patchGuard.OrignPtr)
	funImpl.code = (uintptr)(patchGuard.OrignPtr)
	fmt.Print(funImpl)

	fmt.Print(value1)
	fmt.Println("\nbefore")
	func1.Interface().(func(a int) int)(2)

	fmt.Printf("ok")

	time2(6)
}

type makeFuncImpl struct {
	code  uintptr
	stack unsafe.Pointer
	typ   unsafe.Pointer
	fn    func([]reflect.Value) []reflect.Value
}

const (
	kindMask        = (1 << 5) - 1
	flagMethodShift = 9
)

type flag_t uintptr

type value struct {
	typ *rtype
	ptr unsafe.Pointer
	flag_t
}

type rtype struct {
	size          uintptr
	ptrdata       uintptr
	hash          uint32         // hash of type; avoids computation in hash tables
	_             uint8          // unused/padding
	align         uint8          // alignment of variable with this type
	fieldAlign    uint8          // alignment of struct field with this type
	kind          uint8          // enumeration for C
	alg           uintptr        // algorithm table
	gcdata        uintptr        // garbage collection data
	_string       uintptr        // string form; unnecessary but undeniably useful
	*uncommonType                // (relatively) uncommon fields
	ptrToThis     *rtype         // type for pointer to this type, if used in binary or has methods
	zero          unsafe.Pointer // pointer to zero value
}

type uncommonType struct {
	name    *string  // name of type
	pkgPath *string  // import path; nil for built-in types like int, string
	methods []method // methods associated with type
}

// Method on non-interface type
type method struct {
	name    *string        // name of method
	pkgPath *string        // nil for exported Names; otherwise import path
	mtyp    *rtype         // method type (without receiver)
	typ     *rtype         // .(*FuncType) underneath (with receiver)
	ifn     unsafe.Pointer // fn used in interface call (one-word receiver)
	tfn     unsafe.Pointer // fn used for normal method call
}

// imethod represents a method on an interface type
type imethod struct {
	name    *string // name of method
	pkgPath *string // nil for exported Names; otherwise import path
	typ     *rtype  // .(*FuncType) underneath
}

type interfaceType struct {
	rtype   `reflect:"interface"`
	methods []imethod // sorted by hash
}

func (t *rtype) Kind() reflect.Kind { return reflect.Kind(t.kind & kindMask) }
