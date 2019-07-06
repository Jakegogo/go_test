package main

import (
	asp "github.com/Jakegogo/aspectgo/aspect"
	"regexp"
)

// RedisNewAspect apspect for Cmder.readReply calling
type RedisNewAspect struct {
}

// Pointcut Executed on compilation-time
func (a *RedisNewAspect) Pointcut() asp.Pointcut {
	pkg := regexp.QuoteMeta("(github.com/go-redis/redis.Cmder).readReply")
	return asp.NewCallPointcutFromRegexp(pkg)
}

// OnRecord calling when record
func (a *RedisNewAspect) Advice(ctx asp.Context) []interface{} {
	return nil
}

