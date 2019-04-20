package sexpr

import (
	"LogReplay-sdk-go/client/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

type S1 struct {
	Name string
}

func TestGbKEncode(t *testing.T) {
	val := S1{Name: "中国"}
	var b *bytes.Buffer = &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	if err := encoder.Encode(val); err != nil {
		fmt.Printf("Failed to marshal fields to JSON, %v", err)
	}
	utils.RedisClient.Set("test_key", b.Bytes(), utils.REDIS_LIVE_TIME)
	fmt.Println((string)(b.Bytes()))

	fmt.Println(utils.RedisClient.Get("test_key").String())

}
