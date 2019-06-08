package main

import (
	"fmt"
	asp "github.com/Jakegogo/aspectgo/aspect"
	async_asp "github.com/Jakegogo/aspectgo/aspect/async"
	"github.com/tylerb/gls"
	"go_test/base"
	"regexp"
)

//针对获取当前机器时间的切面
//------------------------
// UnixTimeAspect implements interface asp.Aspect
type GlsAspect struct {
	async_asp.DefaultAsyncGoStmtAspect
}

// -------------------------
// for context pass through go async statement
// on context get implements
// c := OnContextGet()
func (a *GlsAspect) OnContextGet() interface{} {
	return gls.Get("key")
}

// on context set implements
// OnContextSet(c)
func (a *GlsAspect) OnContextSet(ctx interface{}) interface{} {
	gls.Set("key", ctx)
	return nil
}

// on context clear implements
// OnContextClear(c)
func (a *GlsAspect) OnContextClear(ctx interface{}) {
	gls.Set("key", nil)
}

// default global var pointcut
// eg:
//------------------------
type GlobalVarAspect struct {
	async_asp.DefaultGlobalVarAspect
	base.BaseContextAspect
}

// Executed ONLY on runtime
func (a *GlobalVarAspect) Advice(ctx asp.Context) []interface{} {
	fmt.Println("visit global var")
	return ctx.Call(ctx.Args())
}

func (a *GlobalVarAspect) OnRecord(ctx asp.Context) []interface{} {
	return nil
}

type TestAspect struct {
}

// Executed on compilation-time
// Must be override
func (a *TestAspect) Pointcut() asp.Pointcut {
	pkg := regexp.QuoteMeta("go_test/aspect.Call")
	return asp.NewCallPointcutFromRegexp(pkg)
}

// Executed ONLY on runtime
func (a *TestAspect) Advice(ctx asp.Context) []interface{} {
	fmt.Println("visit global var")
	return ctx.Call(ctx.Args())
}

//type TestAspect1 struct {
//}
//
//// Executed on compilation-time
//// Must be override
//func (a *TestAspect1) Pointcut() asp.Pointcut {
//	pkg := regexp.QuoteMeta("(*go_test/another.GlobalVar1).Call")
//	return asp.NewCallPointcutFromRegexp(pkg)
//}
//
//
//// Executed ONLY on runtime
//func (a *TestAspect1) Advice(ctx asp.Context) []interface{} {
//	fmt.Println("visit global var")
//	return ctx.Call(ctx.Args())
//}
