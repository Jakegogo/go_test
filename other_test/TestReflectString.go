package main

import (
	"fmt"
	"reflect"
)

type GlobalVar struct {
	Key string // 不支持私有属性, 私有属性无法序列化
}

var GlobalVarIns *GlobalVar = &GlobalVar{
	Key: "true",
}

func main() {
	fmt.Println(reflect.TypeOf(GlobalVarIns).String())
}
