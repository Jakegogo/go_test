package aspect

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
	"sync"
	"testing"
)

var lock *sync.RWMutex = new(sync.RWMutex)

var lock1 **sync.RWMutex = &lock

type RedisApi struct {
	client *redis.Client
	ip     *string
	port   *int
}

var g_tongcheng_redis *RedisApi

func TestGoVariable(t *testing.T) {

	for i := 1; i <= 10; i++ {
		var n = i
		var finish bool
		go func(n1 int) {
			if n1 == 2 {
				Finish(&finish)
			}
			fmt.Printf("n %d isFinish %t\n", n, finish)
		}(n)
	}

}

func Finish(i *bool) {
	*i = true
}

func TestVarType(t *testing.T) {
	fmt.Println(GetElemType(lock1))
}

func GetElemType(val interface{}) string {
	return elemType(reflect.TypeOf(val), 0)
}

func elemType(typ reflect.Type, stack int) string {
	if stack > 10 {
		return typ.String()
	}
	if typ.Kind() == reflect.Ptr {
		return elemType(typ.Elem(), stack+1)
	}
	return typ.String()
}

var typeMap map[string]map[string]bool = make(map[string]map[string]bool, 100)

func ContainsType(val interface{}, typ string) bool {
	valTyp := GetElemType(val)
	if typs, ok := typeMap[valTyp]; ok {
		if contains, ok1 := typs[typ]; ok1 && contains {
			return true
		}
	}

	result := findType("", reflect.TypeOf(val), typ)
	if _, ok := typeMap[valTyp]; !ok {
		typeMap[valTyp] = make(map[string]bool, 10)
	}

	typeMap[valTyp][typ] = result
	return result
}

//!+display
func findType(path string, v reflect.Type, typ string) bool {
	switch v.Kind() {
	case reflect.Invalid:
		return false
	case reflect.Slice, reflect.Array:
		return findType(fmt.Sprintf("[]%s", path), v.Elem(), typ)
	case reflect.Struct:
		if v.String() == typ {
			return true
		}
		var result bool
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Field(i).Name)
			result = result || findType(fieldPath, v.Field(i).Type, typ)
		}
		return result
	case reflect.Map:
		result := findType(fmt.Sprintf("%s:map[%s]", path,
			v.Key()), v.Key(), typ)
		return result || findType(fmt.Sprintf("%s:map[][%s]", path,
			v.Elem()), v.Key(), typ)
	case reflect.Ptr:
		return findType(fmt.Sprintf("(*%s.%s)", path, v.Elem()), v.Elem(), typ)
	case reflect.Interface:
		return v.String() == typ
	default: // basic types, channels, funcs
		return false
	}
	return false
}

func TestContainType(t *testing.T) {
	fmt.Println(ContainsType(g_tongcheng_redis, "redis.Client"))
}
