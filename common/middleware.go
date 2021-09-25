package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {

	currDir, _ := os.Getwd()
	logFilePath := path.Join(currDir, "logs")
	logFileName := "school"

	fileName := path.Join(logFilePath, logFileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	//logWriter, err := rotatelogs.New(
	//	// 分割后的文件名称
	//	fileName + ".%Y%m%d.log",
	//
	//	// 生成软链，指向最新日志文件
	//	rotatelogs.WithLinkName(fileName),
	//
	//	// 设置最大保存时间(7天)
	//	rotatelogs.WithMaxAge(60*24*time.Hour),
	//
	//	// 设置日志切割时间间隔(1天)
	//	rotatelogs.WithRotationTime(24*time.Hour),
	//)
	//
	//writeMap := lfshook.WriterMap{
	//	logrus.InfoLevel:  logWriter,
	//	logrus.FatalLevel: logWriter,
	//	logrus.DebugLevel: logWriter,
	//	logrus.WarnLevel:  logWriter,
	//	logrus.ErrorLevel: logWriter,
	//	logrus.PanicLevel: logWriter,
	//}
	//
	//lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
	//	TimestampFormat:"2006-01-02 15:04:05",
	//})
	//
	//// 新增 Hook
	//logger.AddHook(lfHook)
	//
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
		}).Info()

	}

}