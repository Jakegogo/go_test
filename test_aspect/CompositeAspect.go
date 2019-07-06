package main

import (
	asp "github.com/Jakegogo/aspectgo/aspect"
)

type CompositeAspect struct {
}

// Executed on compilation-time
func (a *CompositeAspect) Pointcut() asp.Pointcut {
	return asp.NewCallPointcutFromRegexp(`\(*go_test/aspect3\.inner\)\.Call`)
}

// Executed ONLY on runtime
func (a *CompositeAspect) Advice(ctx asp.Context) []interface{} {
	return ctx.Call(ctx.Args())
}