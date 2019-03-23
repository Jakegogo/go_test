package typeassert

import "reflect"

var TypeRegistry = make(map[string]reflect.Type)

func init() {
	typ := reflect.TypeOf(S1{})
	key1 := typ.PkgPath() + "." + typ.Name()
	TypeRegistry[key1] = typ
}
