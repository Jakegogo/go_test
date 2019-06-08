package another

import "fmt"

type GlobalVar1 struct {
	KEY1 string
}

func (g *GlobalVar1) Call() {
	fmt.Println("GlobalVar1.Call")
}
