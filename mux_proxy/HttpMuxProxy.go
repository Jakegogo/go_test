package sexpr

import (
	"fmt"
	echoutil "go_test/echo"
	"net/http"
	"reflect"
	"unsafe"
)

/**
 * 上下文
 */
type ProxyContext struct {
	res  http.ResponseWriter
	req  *http.Request
	next func(http.ResponseWriter, *http.Request)
}

/**
 * 拦截器函数
 */
type Intercepter func(context *ProxyContext) error

/**
 * 添加httpServer的切面
 * 添加到mux上
 * @param mux
 * @param intercepter 拦截器函数
 */
func addIntercepter(
	mux *echoutil.MyServMux,
	intercepter func(context *ProxyContext) error) {
	mField := reflect.ValueOf(&mux.ServeMux).Elem().FieldByName("m")
	// 转换成自定义map指针类型,解决无法修改私有属性的问题
	var muxMap *map[string]muxEntry = (*map[string]muxEntry)(unsafe.Pointer(mField.Addr().Pointer()))
	for key, value := range *muxMap {
		handler1 := createHandler(value.h, intercepter)
		(*muxMap)[key] = muxEntry{h: handler1, pattern: "proxy"}
	}
}

type muxEntry struct {
	h       http.Handler
	pattern string
}

/**
 * 代理Handler类
 */
type ProxyHandler struct {
	origin      http.Handler
	name        string
	intercepter Intercepter
}

/**
 * http请求切面方法
 */
func (h *ProxyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// 请求处理之前
	fmt.Println("proxy handler before call")
	if h.intercepter != nil {
		next := func(innerRes http.ResponseWriter, innerReq *http.Request) {
			// 调用被代理的ServeHTTP方法
			h.origin.ServeHTTP(innerRes, innerReq)
		}
		// 触发拦截器
		context := &ProxyContext{
			res:  res,
			req:  req,
			next: next,
		}
		h.intercepter(context)
	}

	// 请求处理之后
	fmt.Println("proxy handler after call")
}

/**
 * 创建代理Handler
 */
func createHandler(h http.Handler,
	intercepter func(context *ProxyContext) error) http.Handler {
	return &ProxyHandler{origin: h, name: "proxy", intercepter: intercepter}
}
