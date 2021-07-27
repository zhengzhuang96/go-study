// Author: zhengzhuang
// Date: 2021-07-27 10:45:03
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-27 10:45:16
// Description: 日志中间件
// FilePath: /go-study/gin-demo/middleware/logger.go
package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var AppLog *logrus.Logger
var WebLog *logrus.Logger

func Setup() {
	initAppLog()
	initWebLog()
}

// 初始化AppLog
func initAppLog() {
	logFileName := "app.log"
	AppLog = initLog(logFileName)
}

// 初始化WebLog
func initWebLog() {
	logFileName := "web.log"
	WebLog = initLog(logFileName)
}

// 初始化日志句柄
func initLog(logFileName string) *logrus.Logger {

	logName := fmt.Sprintf("%s%s", "logs/", logFileName)

	var src *os.File
	var err error

	// 判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err = os.Stat(logName); os.IsNotExist(err) {
		src, err = os.Create(logName)
	} else {
		src, err = os.OpenFile(logName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 设置日志级别
	logger.Level = logrus.InfoLevel

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		logName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		fmt.Println("err", err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return logger
}

// Gin中间件函数，记录请求日志
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		WebLog.WithFields(logrus.Fields{
			"http_status": statusCode,
			"total_time":  latencyTime,
			"ip":          clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}).Info()
	}
}
