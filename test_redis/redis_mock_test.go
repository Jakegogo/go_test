package test_redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

var RedisClient = redis.NewClusterClient(&redis.ClusterOptions{
	WriteTimeout: time.Millisecond * 800,
	PoolSize: 100,
	Addrs: []string{"9.22.59.140:7000", "9.22.59.140:7001", "9.22.26.217:7002","9.22.26.217:7003", "9.22.26.200:7004", "9.22.26.200:7005"},
})

func TestRedisRead(t *testing.T) {
	val := RedisClient.Get("weishi_qq_svr_l5_cache_480513_131072").String()
	fmt.Println(val)
}