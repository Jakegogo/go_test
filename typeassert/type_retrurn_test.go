package typeassert

import (
	"reflect"
	"testing"
)

func TestTypeReturn(t *testing.T) {
	EC(nil)
}

func m(a interface{}) map[string]bool {
	return reflect.ValueOf(a).Interface().(map[string]bool)
}

func a(a interface{}) []string {
	return reflect.ValueOf(a).Interface().([]string)
}

type S struct {
}

func s(a interface{}) S {
	return reflect.ValueOf(a).Interface().(S)
}

type I interface {
}

func i(a interface{}) I {
	return reflect.ValueOf(a).Interface()
}

type EI interface {
	Error() string
}

func ei(a interface{}) EI {
	return reflect.ValueOf(a).Interface().(EI)
}

func e(a interface{}) error {
	return reflect.ValueOf(a).Interface().(error)
}

func EC(a interface{}) (error, EI) {
	return func(a interface{}) error {
			if a == nil {
				return nil
			}
			return reflect.ValueOf(a).Interface().(error)
		}(a), func(a interface{}) EI {
			if a == nil {
				return nil
			}
			return reflect.ValueOf(a).Interface().(EI)
		}(a)
}
