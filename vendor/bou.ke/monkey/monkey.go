package monkey // import "bou.ke/monkey"

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

// patch is an applied patch
// needed to undo a patch
type patch struct {
	originalBytes []byte
	replacement   *reflect.Value
}

var (
	lock = sync.Mutex{}

	patches = make(map[uintptr]patch)
)

type value struct {
	_   uintptr
	ptr unsafe.Pointer
}

func getPtr(v reflect.Value) unsafe.Pointer {
	return (*value)(unsafe.Pointer(&v)).ptr
}

func getInfirectPtr(v reflect.Value) unsafe.Pointer {
	return (*value)(unsafe.Pointer(&v)).ptr
}

type PatchGuard struct {
	target       reflect.Value
	replacement  reflect.Value
	OrignPtr     unsafe.Pointer
	OrignBytes   *[]byte
	OrignUintptr uintptr
}

func (g *PatchGuard) Unpatch() {
	unpatchValue(g.target)
}

func (g *PatchGuard) Restore() {
	patchValue(g.target, g.replacement, reflect.ValueOf(nil))
}
func Patch(target, replacement interface{}) *PatchGuard {
	return Patch1(target, replacement, nil)
}

// Patch replaces a function with another
func Patch1(target, replacement, placehlder interface{}) *PatchGuard {
	t := reflect.ValueOf(target)
	r := reflect.ValueOf(replacement)
	p := reflect.ValueOf(placehlder)
	orignPtr, originJumpData, OrignUintptr := patchValue(t, r, p)
	return &PatchGuard{t, r, orignPtr, originJumpData, OrignUintptr}
}

func PatchWithJump(target interface{}, jumpData *[]byte) *PatchGuard {
	t := reflect.ValueOf(target)
	orignPtr, orignJumpData, OrignUintptr := patchValueWithJump(t, jumpData)

	return &PatchGuard{t, reflect.Zero(t.Type()), orignPtr, orignJumpData, OrignUintptr}
}

// PatchInstanceMethod replaces an instance method methodName for the type target with replacement
// Replacement should expect the receiver (of type target) as the first argument
func PatchInstanceMethod(target reflect.Type, methodName string, replacement interface{}) *PatchGuard {
	m, ok := target.MethodByName(methodName)
	if !ok {
		panic(fmt.Sprintf("unknown method %s", methodName))
	}
	r := reflect.ValueOf(replacement)
	patchValue(m.Func, r, reflect.ValueOf(nil))

	return &PatchGuard{m.Func, r, nil, nil, 0}
}

func patchValue(target, replacement, placehlder reflect.Value) (unsafe.Pointer, *[]byte, uintptr) {
	lock.Lock()
	defer lock.Unlock()

	if target.Kind() != reflect.Func {
		panic("target has to be a Func")
	}

	if replacement.Kind() != reflect.Func {
		panic("replacement has to be a Func")
	}

	if target.Type() != replacement.Type() {
		panic(fmt.Sprintf("target and replacement have to have the same type %s != %s", target.Type(), replacement.Type()))
	}

	if patch, ok := patches[target.Pointer()]; ok {
		unpatch(target.Pointer(), patch)
	}

	p := unsafe.Pointer(target.Pointer())
	p1 := *(*unsafe.Pointer)(p)
	ptr := uintptr(p1)
	fmt.Println(" ptr is:", ptr)

	bytes, orignFunc, targetuinptr := replaceFunction(target.Pointer(), (uintptr)(getPtr(replacement)), (uintptr)(getPtr(target)), placehlder.Pointer()+20)
	patches[target.Pointer()] = patch{*bytes, &replacement}

	//targetInterface := target.Interface()
	//targetuinptr := *(*uintptr)(unsafe.Pointer(&targetInterface))
	//targetuinptr := (uintptr)(target.Pointer())
	//18479480
	//originJumpData := callToFunctionValue(targetuinptr)

	//mprotectCrossPage((uintptr)(unsafe.Pointer(&originJumpData[0])), len(originJumpData), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)

	//fmt.Printf("target uinptr is %d\n", (uintptr)(unsafe.Pointer(&originJumpData[0])))

	fmt.Printf("target uinptr is %d\n", targetuinptr)
	// ok
	//return orignFunc, &originJumpData, targetuinptr
	// not ok
	//return orignFunc, &originJumpData, (uintptr)(unsafe.Pointer(&originJumpData[0]))
	return orignFunc, bytes, targetuinptr
}

func patchValueWithJump(target reflect.Value, jumpData *[]byte) (unsafe.Pointer, *[]byte, uintptr) {
	lock.Lock()
	defer lock.Unlock()

	if patch, ok := patches[target.Pointer()]; ok {
		unpatch(target.Pointer(), patch)
	}

	bytes, orignFunc, originPtr := replaceFunctionWithJump(target.Pointer(), (uintptr)(getPtr(target)), jumpData)
	patches[target.Pointer()] = patch{bytes, nil}

	return orignFunc, &bytes, originPtr
}

// Unpatch removes any monkey patches on target
// returns whether target was patched in the first place
func Unpatch(target interface{}) bool {
	return unpatchValue(reflect.ValueOf(target))
}

// UnpatchInstanceMethod removes the patch on methodName of the target
// returns whether it was patched in the first place
func UnpatchInstanceMethod(target reflect.Type, methodName string) bool {
	m, ok := target.MethodByName(methodName)
	if !ok {
		panic(fmt.Sprintf("unknown method %s", methodName))
	}
	return unpatchValue(m.Func)
}

// UnpatchAll removes all applied monkeypatches
func UnpatchAll() {
	lock.Lock()
	defer lock.Unlock()
	for target, p := range patches {
		unpatch(target, p)
		delete(patches, target)
	}
}

// Unpatch removes a monkeypatch from the specified function
// returns whether the function was patched in the first place
func unpatchValue(target reflect.Value) bool {
	lock.Lock()
	defer lock.Unlock()
	patch, ok := patches[target.Pointer()]
	if !ok {
		return false
	}
	unpatch(target.Pointer(), patch)
	delete(patches, target.Pointer())
	return true
}

func unpatch(target uintptr, p patch) {
	copyToLocation(target, p.originalBytes)
}
