package sexpr

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go_test/xerrs"
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: false})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func NewRecordHttpIntercepter() func(context *ProxyContext) error {
	return func(context *ProxyContext) error {
		return RecordHttpIntercepter(context.res, context.req, context.next)
	}
}

/**
 * 请求录制拦截器
 */
func RecordHttpIntercepter(res http.ResponseWriter, req *http.Request, next func(http.ResponseWriter, *http.Request)) error {
	var body []byte

	// 读取body内容
	if bodyReader := req.Body; bodyReader != nil {
		body, _ = ioutil.ReadAll(bodyReader)
	}

	var responseBody bytes.Buffer
	// 使用代理Context包装,从而可以捕获到response.Write写入的内容
	wrapWriter := NewRecordResponseWriter(&res, &responseBody)

	var err error
	// 执行下一步
	err = xerrs.Try(func() {
		next(wrapWriter, req)
	})

	if err != nil {
		xerrs.Print(err)
	}

	// 输出到日志
	log.WithFields(log.Fields{
		"url":            req.URL,
		"requestHeader":  req.Header,
		"requestBody":    body,
		"responseHeader": res.Header(),
		"responseBody":   responseBody.String(),
		"status":         wrapWriter.GetStatusCode(),
		"error":          xerrs.Details(err, 30),
	}).Info("request content")

	fmt.Println(responseBody.String())

	return err
}
