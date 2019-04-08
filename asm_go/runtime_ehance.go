package main

import (
	"LogReplay-sdk-go/client/utils"
	"fmt"
	forcexport "go_test/forcexport"
)

func main() {
	utils.RedisClient.ZAdd("")

	ptr, _ := forcexport.FindFuncWithName("github.com/PinkDahlia/redis.(*cmdable).zAdd")
	fmt.Printf("Found pointer 0x%x", ptr)
}
