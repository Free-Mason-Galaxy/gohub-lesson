// Package middlewares
// descr recovery中间件
// author fm
// date 2022/11/16 14:50
package middlewares

import (
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gohub-lesson/pkg/logger"
	"gohub-lesson/pkg/response"
	"gorm.io/gorm"
)

// Recovery 使用 zap.Error() 来记录 panic 和 call stack
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				tx, _ := ctx.Get("transaction")

				if tx != nil {
					tx.(*gorm.DB).Rollback()
				}

				// 获取用户的请求信息
				httpRequest, _ := httputil.DumpRequest(ctx.Request, true)

				// 链接中断，客户端中断连接为正常行为，不需要记录堆栈信息
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "broken pipe") || strings.Contains(errStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 链接中断的情况
				if brokenPipe {
					logger.Error(ctx.Request.URL.Path,
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					_ = ctx.Error(err.(error))
					ctx.Abort()
					// 链接已断开，无法写状态码
					return
				}

				// 如果不是链接中断，就开始记录堆栈信息
				logger.Error("recovery from panic",
					zap.Time("time", time.Now()),               // 记录时间
					zap.Any("error", err),                      // 记录错误信息
					zap.String("request", string(httpRequest)), // 请求信息
					zap.Stack("stacktrace"),                    // 调用堆栈信息
				)

				// 返回 500 状态码
				response.Abort500(ctx)
			}
		}()
		ctx.Next()
	}
}
