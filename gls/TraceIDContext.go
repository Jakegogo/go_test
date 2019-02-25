package gls

import (
	"LogReplay-sdk-go/client/utils"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tylerb/gls"
	"strconv"
	"time"
)

const TRACE_KEY string = "_logreplay_thread_local_trace_id_"

var auto_inc *utils.AutoInc = utils.NewAutoInc(0, 1)

/**
 * TraceID上下文
 */
type TraceIDContext struct {
	uuid string // 唯一traceId
}

/**
 * 获取Trace上下文
 * @return *TraceIDContext 上下文
 */
func GetContext() *TraceIDContext {
	var curContext = gls.Get(TRACE_KEY)
	if curContext != nil {
		return curContext.(*TraceIDContext)
	}

	var traceID string = ""
	u2, err := uuid.NewV4()
	if err != nil {
		// 生成失败则用毫秒数代替
		rand := time.Now().Nanosecond()*10000 + auto_inc.Id()
		traceID = strconv.Itoa(rand)
	}

	traceID = u2.String()
	context := &TraceIDContext{
		uuid: traceID,
	}
	gls.Set(TRACE_KEY, context)
	return context
}

/**
 * 设置上下文,用于多线程间传递
 * @param *TraceIDContext 上下文
 */
func SetContext(context *TraceIDContext) {
	var curContext = gls.Get(TRACE_KEY)
	if curContext != nil {
		log.WithFields(log.Fields{
			"uuid": context.uuid,
		}).Warn("trace context not empty..")
	}

	gls.Set(TRACE_KEY, context)
}

/**
 * 获取TraceID
 * @return string 上下文ID
 */
func GetTraceID() string {
	context := GetContext()
	return context.uuid
}

/**
 * 清理trace上下文
 */
func Clear() {
	gls.Set(TRACE_KEY, nil)
}
