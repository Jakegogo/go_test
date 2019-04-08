package main

import (
	"LogReplay-sdk-go/client/utils"
	"fmt"
	"github.com/PinkDahlia/redis"
	forcexport "go_test/forcexport"
	"reflect"
	"unsafe"
)

func main() {
	utils.RedisClient.ZAdd("")

	ptr, _ := forcexport.FindFuncWithName("github.com/PinkDahlia/redis.(*cmdable).ZAdd")
	fmt.Printf("Found pointer 0x%x", ptr)

	cmdable := reflect.ValueOf(utils.RedisClient).Elem().FieldByName("cmdable").UnsafeAddr()
	cmdAble := (*Cmdable)(unsafe.Pointer(cmdable))

	// 访问私有类型的方法
	var fn = cmdAble.ZAdd
	fn("w")

	forcexport.CreateFuncForCodePtr(&fn, ptr)
	fn("w")
}

// not work
type Cmdable struct {
	process func(cmd redis.Cmder) error
}

func (c *Cmdable) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	return nil
}

func ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	return nil
}
