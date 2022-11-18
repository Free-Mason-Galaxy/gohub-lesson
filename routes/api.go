// Package routes
// descr 注册路由
// author fm
// date 2022/11/14 16:22
package routes

import (
	"github.com/gin-gonic/gin"
	"gohub-lesson/app/http/controllers/api/v1/auth"
	"gohub-lesson/app/http/controllers/test"
)

func RegisterAPIRoutes(r *gin.Engine) {

	t := r.Group("/")
	{
		testController := new(test.TestController)
		t.Any("/test", testController.Any)
	}

	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			signController := new(auth.SignupController)
			// 判断手机号是否存在
			authGroup.GET("/signup/phone/exist", signController.IsPhoneExist)
			// 判断邮箱是否存在
			authGroup.GET("/signup/email/exist", signController.IsEmailExist)

			sendVerifyCodeController := new(auth.SendVerifyCodeController)
			// 获取图片验证码
			authGroup.POST("/verify-codes/captcha", sendVerifyCodeController.ShowCaptcha)
			// 发送手机验证码
			authGroup.POST("/verify-codes/phone", sendVerifyCodeController.SendUsingPhone)
			// 发送邮件
			authGroup.POST("/verify-codes/email", sendVerifyCodeController.SendEmail)
		}
	}
}
