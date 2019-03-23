package sexpr

import (
	"reflect"
	"unsafe"
)

/**
 * 设置非导出结构体字段的值
 * 1.可适用目标struct为引用类型的参数
 * 2.如果目标struct为非引用类型参数:
 *  a,如果目标字段为引用型可以进行设值;
 *  b,如果目标字段为非引用型,则无法在原对象上修改。会创建新的对象为字段设值,并返回新的对象
 * @param structV 目标结构体
 * @param name 字段名
 * @param v 字段值
 * @return interfacetest.(*Type)
 */
func SetUnexportField(structV interface{}, name string, v interface{}) (newStruct interface{}) {
	return setUnexport(reflect.ValueOf(structV), name, reflect.ValueOf(v))
}

/**
 * 设置非导出结构体的字段值
 * @param structV 目标结构体
 * @param name 字段名
 * @param v 字段值
 * @return interfacetest.(*Type)
 */
func setUnexport(structV reflect.Value, name string, v reflect.Value) (newStruct interface{}) {
	var fieldV reflect.Value
	var rs2 reflect.Value
	// 引用和接口类型的结构体可以直接寻址
	if structV.Kind() == reflect.Ptr || structV.Kind() == reflect.Interface {
		emv := structV.Elem()
		fieldV = emv.FieldByName(name)
	} else {
		// 非引用类型或者接口,拷贝一份结构体值进行寻址
		rs2 = reflect.New(structV.Type()).Elem()
		rs2.Set(structV)
		fieldV = rs2.FieldByName(name)
		// Now fieldV can be read.  Setting will succeed but only affects the temporary copy.
	}

	// 指针或者接口类型,需要进行寻址
	if fieldV.Kind() == reflect.Ptr || structV.Kind() == reflect.Interface {
		fieldV = fieldV.Elem()
	}

	// 是否为导出类型
	if fieldV.CanSet() {
		fieldV.Set(v)
	} else {
		fieldV = reflect.NewAt(fieldV.Type(), unsafe.Pointer(fieldV.UnsafeAddr())).Elem()
		// Now fieldV can be read and set.
		fieldV.Set(v)
	}

	if rs2.IsValid() {
		return rs2.Addr().Interface()
	} else {
		return structV.Interface()
	}

}
