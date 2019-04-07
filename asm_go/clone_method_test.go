package main

import "reflect"

//called while building the []reflect.Method to pass to NewTypeOf
func cloneMethod(onto reflect.Type, rcvr string, m reflect.Method) reflect.Method {
	in := make([]reflect.Type, m.Type.NumIn())
	in[0] = onto //XXX we need our receiver type but this is the type of the underlying struct not the actual type
	for i := 1; i < m.Type.NumIn(); i++ {
		in[i] = m.Type.In(i)
	}

	out := make([]reflect.Type, m.Type.NumOut())
	for i := 0; i < m.Type.NumOut(); i++ {
		out[i] = m.Type.Out(i)
	}

	typ := reflect.FuncOf(in, out, m.Type.IsVariadic())
	fn := reflect.MakeFunc(typ, func(args []reflect.Value) []reflect.Value {
		r := args[0].FieldByName(rcvr)
		args = append([]reflect.Value{r}, args[1:]...)
		if m.Type.IsVariadic() {
			return r.CallSlice(args[1:])
		} else {
			return r.Call(args[1:])
		}
	})

	return reflect.Method{
		Name:    m.Name,
		PkgPath: m.PkgPath,
		Type:    typ,
		Func:    fn,
	}
}
