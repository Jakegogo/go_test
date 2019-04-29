package main

// go tool compile -S asm_iface.go >> asm_iface.s
type Mather interface {
	Add(a, b int32) int32
	Sub(a, b int64) int64
}

type Adder struct{ id int32 }

//go:noinline
func (adder Adder) Add(a, b int32) int32 { return a + b }

//go:noinline
func (adder Adder) Sub(a, b int64) int64 { return a - b }

//func main() {
//	m := Mather(Adder{id: 6754})
//	m.Add(10, 32)
//}
