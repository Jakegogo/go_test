package main

import (
	"fmt"
	"regexp"
	"runtime"

	asp "github.com/AkihiroSuda/aspectgo/aspect"
)

//true	to record
//false	to playback
var recordOrPlayback bool

//var redisClient *redis.Client

func init() {
	recordOrPlayback = true
}

func trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// ExampleAspect implements interface asp.Aspect
type ExampleAspect struct {
}

// Executed on compilation-time
func (a *ExampleAspect) Pointcut() asp.Pointcut {
	pkg := regexp.QuoteMeta("xorm_with_aspect")
	s := pkg + ".*"
	return asp.NewCallPointcutFromRegexp(s)
}

// Executed ONLY on runtime
func (a *ExampleAspect) Advice(ctx asp.Context) []interface{} {
	args := ctx.Args()
	key := trace() + "_" + fmt.Sprint(args)
	fmt.Printf("record key : %v\n", key)
	if recordOrPlayback {
		res := ctx.Call(args)

		//record into redis
		//err := redisClient.Set(key, res[0], 0).Err()
		//if err != nil {
		//	fmt.Printf("Error occured when set to redis : %v\n", err)
		//}

		return res
	} else {
		fmt.Println("got from redis")
		//val := redisClient.Get(key).Val()

		return []interface{}{val}
	}
}
