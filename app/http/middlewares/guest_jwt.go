// Package middlewares
// descr Guest 中间件
// author fm
// date 2022/11/21 14:48
package middlewares

import (
	"github.com/gin-gonic/gin"
	"gohub-lesson/pkg/helpers"
	"gohub-lesson/pkg/jwt"
	"gohub-lesson/pkg/response"
)

// GuestJWT 强制使用游客身份访问
func GuestJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if len(ctx.GetHeader("Authorization")) > 0 {

			// 解析 token 成功，说明登录成功了
			_, err := jwt.NewJWT().ParseToken(ctx)

			if !helpers.IsError(err) {
				response.Unauthorized(ctx, "请使用游客身份访问")
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}
