// Package middlewares
// descr 储放系统日志中间件
// author fm
// date 2022/11/16 10:31
package middlewares

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gohub-lesson/pkg/helpers"
	"gohub-lesson/pkg/logger"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取 response 内容
		w := &responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}

		ctx.Writer = w

		// 获取数据
		var requestBody []byte
		if ctx.Request.Body != nil {
			// ctx.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = io.ReadAll(ctx.Request.Body)
			// 读取后，重新赋值 ctx.Request.Body ，以供后续的其他操作
			ctx.Request.Body = io.NopCloser(bytes.NewReader(requestBody))
		}

		// 设置开始时间
		start := time.Now()
		ctx.Next()

		// 开始记录日志的逻辑
		cost := time.Since(start)
		responseStatus := ctx.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", responseStatus),
			zap.String("request", ctx.Request.Method+" "+ctx.Request.URL.String()),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}

		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
			// 请求的内容
			logFields = append(logFields, zap.String("request-body", string(requestBody)))

			// 响应的内容
			logFields = append(logFields, zap.String("response-body", w.body.String()))
		}

		if responseStatus > 400 && responseStatus <= 499 {
			// 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
			logger.Warn("HTTP Warning "+cast.ToString(responseStatus), logFields...)
		} else if responseStatus >= 500 && responseStatus <= 599 {
			// 除了内部错误，记录 error
			logger.Error("HTTP Error "+cast.ToString(responseStatus), logFields...)
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}
	}
}
