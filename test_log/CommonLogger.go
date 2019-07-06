package test_log

import (
	"LogReplay-sdk-go/client/plugins/registry"
	"context"
	"github.com/heetch/confita/backend/env"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unsafe"
)


/**
 * 日志存放路径
 */
const LOGGER_PATH = "/dockerdata/"

// 录制回放存储的日志文件
const RecordLoggerFile = "logreplay_record.log"

// 异常日志文件
const ExceptionLoggerFile = "logreplay_error.log"

// 调试日志文件
const DebugLoggerFile = "logreplay_debug.log"

var defaultFormater = &log.TextFormatter{
	TimestampFormat : "2006-01-02 15:04:05.000",
	FieldMap: log.FieldMap{
		"func":  "caller",
	},
	DisableColors : true,
	DisableSorting: true,
}

var defaultJsonFormater = &log.JSONFormatter{
	TimestampFormat : "2006-01-02 15:04:05.000",
	PrettyPrint: false,
	FieldMap: log.FieldMap{
		"func":  "caller",
	},
}

func init() {

	var logFileLocation string = LOGGER_PATH;
	env := env.NewBackend()
	var envBytes, err = env.Get(context.Background(), "LOGREPLAY_LOGGER_PATH")
	if envBytes != nil {
		logFileLocation = *(*string)(unsafe.Pointer(&envBytes))
	}
	// 获取当前目录
	if wd, err := os.Getwd(); err == nil && strings.Contains(wd, string(filepath.Separator)) && "/" != wd {
		logFileLocation = filepath.Join(wd, "logreplay")
	} else {
		// 按应用名分文件夹
		logFileLocation = filepath.Join(logFileLocation, registry.SvrInfo.Name)
	}

	var Logger = log.StandardLogger()
	Logger.SetLevel(log.DebugLevel)
	//Logger.SetReportCaller(true)
	// 设置时间格式
	formatter, ok := Logger.Formatter.(*log.TextFormatter)
	if ok {
		formatter.TimestampFormat = defaultFormater.TimestampFormat
		formatter.DisableSorting = defaultFormater.DisableSorting
		formatter.FieldMap = defaultFormater.FieldMap
	}

	// ---------- 主日志文件 -----------
	// 设置滚动日志
	recordLogPath := filepath.Join(logFileLocation, RecordLoggerFile)
	recordWriter, err := rotatelogs.New(
		recordLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(recordLogPath),
		rotatelogs.WithMaxAge(7 * 24 * time.Hour),
		rotatelogs.WithRotationTime(24 * time.Hour),
	)

	if err != nil {
		log.Errorf("common logger init fail:%s", err)
	}

	// 录制内容日志文件
	Logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			log.InfoLevel:  recordWriter,
		},
		defaultJsonFormater,
	))


	// ---------- 异常日志文件 -----------
	// 设置滚动日志
	exceptionLogPath := filepath.Join(logFileLocation, ExceptionLoggerFile)
	exceptionWriter, err := rotatelogs.New(
		exceptionLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(exceptionLogPath),
		rotatelogs.WithMaxAge(7 * 24 * time.Hour),
		rotatelogs.WithRotationTime(24 * time.Hour),
	)

	if err != nil {
		log.Errorf("common logger init fail:%s", err)
	}

	// 异常日志,用于查看
	// 主日志文件
	Logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			log.FatalLevel: exceptionWriter,
			log.ErrorLevel: exceptionWriter,
			log.WarnLevel: exceptionWriter,
		},
		defaultJsonFormater,
	))


	// ---------- 调试日志文件 -----------
	// 设置滚动日志
	debugLogPath := filepath.Join(logFileLocation, DebugLoggerFile)
	debugWriter, err := rotatelogs.New(
		debugLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(debugLogPath),
		rotatelogs.WithMaxAge(3 * 24 * time.Hour),
		rotatelogs.WithRotationTime(24 * time.Hour),
	)

	if err != nil {
		log.Errorf("common logger init fail:%s", err)
	}

	// 调试日志,用于查看
	// 主日志文件
	Logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			log.TraceLevel: debugWriter,
			log.FatalLevel: debugWriter,
			log.DebugLevel: debugWriter,
			log.WarnLevel: debugWriter,
			log.InfoLevel: debugWriter,
			log.ErrorLevel: debugWriter,
		},
		defaultFormater,
	))

	log.Debugf("logger init success")
	log.SetLevel(log.DebugLevel)
	Logger.SetLevel(log.DebugLevel)
}
