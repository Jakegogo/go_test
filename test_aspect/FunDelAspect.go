package main

import (
	asp "github.com/Jakegogo/aspectgo/aspect"
	async_asp "github.com/Jakegogo/aspectgo/aspect/async"
	"regexp"
)

//针对获取当前机器时间的切面
//------------------------
// UnixTimeAspect implements interface asp.Aspect
type FunDelAspect struct {
	async_asp.DefaultAstInsertStmtAspect
}

// Executed on compilation-time
// Must be override
func (a *FunDelAspect) Pointcut() asp.Pointcut {
	pkg := regexp.QuoteMeta("(*go_test/aspect.Obj).AsyncCall")
	return asp.NewCallPointcutFromRegexp(pkg)
}

// Executed ONLY on runtime
func (a *FunDelAspect) Advice(ctx asp.Context) []interface{} {
	return ctx.Call(ctx.Args())
}

func (a *FunDelAspect) IsAstInsert() bool {
	return true
}
