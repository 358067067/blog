package middleware

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Logger 日志
func Logger() gin.HandlerFunc {
	logger := logrus.New()
	return func(ctx *gin.Context) {
		startTime := time.Now()
		//洋葱模式（火炬传递）
		ctx.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknow"
		}
		statusCode := ctx.Writer.Status()
		clientIp := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		dataSize := ctx.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := ctx.Request.Method
		path := ctx.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
