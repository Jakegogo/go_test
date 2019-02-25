package main

import (
	"bytes"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	muxutil "go_test/mux_proxy"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		c.Response().Write([]byte("ok"))
		return next(c)
	}
}

// Process is the middleware function.
func (s *Stats) Record(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var body []byte

		// 读取body内容
		if bodyReader := c.Request().Body; bodyReader != nil {
			body, _ = ioutil.ReadAll(bodyReader)
		}

		var responseBody bytes.Buffer
		httpResponseWriter := c.Response().Writer
		// 使用代理Context包装,从而可以捕获到response.Write写入的内容
		wrapWriter := muxutil.NewRecordResponseWriter(&httpResponseWriter, &responseBody)
		//c = echoutil.NewRecordContext(c, wrapWriter, c.Echo()) // 此方法替换context无效

		// 替换writer
		c.Response().Writer = wrapWriter

		// 执行下一步
		if err := next(c); err != nil {
			c.Error(err)
		}

		//c.Response().Write([]byte("not ok"))

		// 输出到日志
		log.WithFields(log.Fields{
			"url":            c.Request().URL,
			"requestHeader":  c.Request().Header,
			"requestBody":    body,
			"responseHeader": c.Response().Header(),
			"responseBody":   responseBody.String(),
			"status":         c.Response().Status,
		}).Info("request content")

		return nil
	}
}

func main() {
	e := echo.New()

	// Debug mode
	e.Debug = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//-------------------
	// Custom middleware
	//-------------------
	// Stats
	s := NewStats()

	// 请求拦截器,放到最前面
	e.Use(s.Record)

	e.Use(s.Process)
	e.GET("/stats", s.Handle) // Endpoint to get stats

	// Server header
	e.Use(ServerHeader)

	// Handler
	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
