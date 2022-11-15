// Package routes
// descr 注册路由
// author fm
// date 2022/11/14 16:22
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gohub-lesson/app/http/controllers/api/v1/auth"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"test": "hello world",
			})
		})
		authGroup := v1.Group("/auth")
		{
			signController := new(auth.SignupController)
			// 判断手机号是否存在
			authGroup.GET("/signup/phone/exist", signController.IsPhoneExist)
			// 判断邮箱是否存在
			authGroup.GET("/signup/email/exist", signController.IsEmailExist)
		}
	}
}
