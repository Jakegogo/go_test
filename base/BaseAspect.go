package base

import (
	asp "github.com/Jakegogo/aspectgo/aspect"
)

type BaseContextAspect struct {
	Name             string
	OnRecord         func(ctx asp.Context) []interface{}
	OnReplay         func(ctx asp.Context) []interface{}
	OnRecordGlobally func(ctx asp.Context) []interface{}
	OnReplayGlobally func(ctx asp.Context) []interface{}
}

// Executed ONLY on runtime
func (a *BaseContextAspect) Advice(ctx asp.Context) []interface{} {
	return a.OnRecord(ctx)
}
