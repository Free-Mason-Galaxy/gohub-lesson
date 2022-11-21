// Package routes
// descr 注册路由
// author fm
// date 2022/11/14 16:22
package routes

import (
	"github.com/gin-gonic/gin"
	"gohub-lesson/app/http/controllers/api/v1/auth"
	"gohub-lesson/app/http/controllers/test"
	"gohub-lesson/app/http/middlewares"
	pkgAuth "gohub-lesson/pkg/auth"
	"gohub-lesson/pkg/response"
)

func RegisterAPIRoutes(r *gin.Engine) {

	// statsviz 实时可视化Go Runtime指标
	// r.GET("/debug/statsviz/*filepath", func(context *gin.Context) {
	// 	if context.Param("filepath") == "/ws" {
	// 		statsviz.Ws(context.Writer, context.Request)
	// 		return
	// 	}
	// 	statsviz.IndexAtRoot("/debug/statsviz").ServeHTTP(context.Writer, context.Request)
	// })

	r.GET("/test_auth", middlewares.AuthJWT(), func(ctx *gin.Context) {
		userModel := pkgAuth.CurrentUser(ctx)
		response.Data(ctx, userModel)
	})

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
			// 手机号注册
			authGroup.POST("/signup/using-phone", signController.SignupUsingPhone)
			// 邮箱注册
			authGroup.POST("/signup/using-email", signController.SignupUsingEmail)

			sendVerifyCodeController := new(auth.SendVerifyCodeController)
			// 获取图片验证码
			authGroup.POST("/verify-codes/captcha", sendVerifyCodeController.ShowCaptcha)
			// 发送手机验证码
			authGroup.POST("/verify-codes/phone", sendVerifyCodeController.SendUsingPhone)
			// 发送邮件
			authGroup.POST("/verify-codes/email", sendVerifyCodeController.SendEmail)

			loginController := new(auth.LoginController)
			// 手机号登录
			authGroup.POST("/login/using-phone", loginController.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", loginController.LoginByPassword)
			// 重置 token
			authGroup.POST("/login/refresh-token", loginController.RefreshToken)
		}
	}
}
