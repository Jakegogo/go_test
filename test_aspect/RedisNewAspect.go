package main

import (
	"LogReplay-sdk-go/client/aspect/base"
	asp "github.com/Jakegogo/aspectgo/aspect"
	"regexp"
)

// RedisNewAspect apspect for Cmder.readReply calling
type RedisNewAspect struct {
	base.BaseContextAspect
}

// Pointcut Executed on compilation-time
func (a *RedisNewAspect) Pointcut() asp.Pointcut {
	pkg := regexp.QuoteMeta("(github.com/go-redis/redis/Cmder).readReply")
	return asp.NewCallPointcutFromRegexp(pkg)
}

// OnRecord calling when record
func (a *RedisNewAspect) OnRecord() error {
	return nil
}

// OnReplay calling when replay
func (a *RedisNewAspect) OnReplay() []interface{} {
	return []interface{}{}
}