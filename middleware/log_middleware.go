package middleware

import (
	"go_web/config"
	"log"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 日志记录到文件中

func LoggerMiddleware() gin.HandlerFunc {
	logFilePath := config.Set.Log.FilePath
	logFileName := config.Set.Log.FileName

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	// 写入文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal("日志创建错误", err)
	}
	// 初始化日志对象
	logger := logrus.New()
	// 设置输出
	logger.Out = file
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		c.Next() // 调用控制器函数函数
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方法
		reqMethod := c.Request.Method
		// 请求路径
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 日志格式
		logger.Infof("| %3d | %13v | %s | %s ",
			statusCode,
			latencyTime,
			reqMethod,
			reqUri,
		)
	}
}
