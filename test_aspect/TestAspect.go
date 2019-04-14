package main

import (
	async_asp "github.com/Jakegogo/aspectgo/aspect/async"
	"github.com/tylerb/gls"
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
func (a *GlsAspect) OnContextSet(ctx interface{}) {
	gls.Set("key", ctx)
}
