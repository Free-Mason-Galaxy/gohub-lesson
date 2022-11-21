// Package middlewares
// descr
// author fm
// date 2022/11/21 14:32
package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	userModel "gohub-lesson/app/models/user"
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/helpers"
	"gohub-lesson/pkg/jwt"
	"gohub-lesson/pkg/response"
)

// AuthJWT 解析验证 token
func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParseToken(ctx)

		if helpers.IsError(err) {
			response.Unauthorized(ctx, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		user := userModel.Get(claims.UserID)

		if user.NotExists() {
			response.Unauthorized(ctx, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		ctx.Set("current_user_id", user.GetIdString())
		ctx.Set("current_user_name", user.Name)
		ctx.Set("current_user", user)

		ctx.Next()
	}
}
