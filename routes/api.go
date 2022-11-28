// Package routes
// descr 注册路由
// author fm
// date 2022/11/14 16:22
package routes

import (
	"github.com/gin-gonic/gin"
	v1controller "gohub-lesson/app/http/controllers/api/v1"
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

	r.GET("/test_guest", middlewares.GuestJWT(), func(ctx *gin.Context) {
		// userModel := pkgAuth.CurrentUser(ctx)
		response.Success(ctx)
	})

	t := r.Group("/")
	{
		testController := new(test.TestController)
		t.Any("/test", testController.Any)
	}

	v1 := r.Group("/v1")
	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// ------------------------------ 注册 ------------------------------ //
			signController := new(auth.SignupController)
			// 判断手机号是否存在
			authGroup.GET("/signup/phone/exist",
				middlewares.GuestJWT(), signController.IsPhoneExist)
			// 判断邮箱是否存在
			authGroup.GET("/signup/email/exist",
				middlewares.GuestJWT(), signController.IsEmailExist)
			// 手机号注册
			authGroup.POST("/signup/using-phone",
				middlewares.GuestJWT(),
				middlewares.LimitPerRoute("60-H"),
				signController.SignupUsingPhone)
			// 邮箱注册
			authGroup.POST("/signup/using-email",
				middlewares.GuestJWT(),
				middlewares.LimitPerRoute("60-H"),
				signController.SignupUsingEmail)

			// ---------------------------- 发送验证码 ---------------------------- //
			sendVerifyCodeController := new(auth.SendVerifyCodeController)
			// 获取图片验证码
			authGroup.POST("/verify-codes/captcha",
				middlewares.LimitPerRoute("50-H"),
				sendVerifyCodeController.ShowCaptcha)
			// 发送手机验证码
			authGroup.POST("/verify-codes/phone",
				middlewares.LimitPerRoute("20-H"),
				sendVerifyCodeController.SendUsingPhone)
			// 发送邮件
			authGroup.POST("/verify-codes/email",
				middlewares.LimitPerRoute("20-H"),
				sendVerifyCodeController.SendEmail)

			// ------------------------------ 登录 ------------------------------ //
			loginController := new(auth.LoginController)
			// 手机号登录
			authGroup.POST("/login/using-phone",
				middlewares.GuestJWT(),
				loginController.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password",
				middlewares.GuestJWT(),
				loginController.LoginByPassword)
			// 重置 token
			authGroup.POST("/login/refresh-token",
				middlewares.AuthJWT(),
				loginController.RefreshToken)

			// ----------------------------- 重置密码 ------------------------------ //
			pwdController := new(auth.PasswordController)
			// 使用手机号重置密码
			authGroup.POST("/password-reset/using-phone",
				middlewares.GuestJWT(),
				pwdController.ResetByPhone)
			// 使用邮箱重置密码
			authGroup.POST("/password-reset/using-email",
				middlewares.GuestJWT(),
				pwdController.ResetByEmail)

		}

		usersController := new(v1controller.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), usersController.CurrentUser)
		usersGroup := v1.Group("/users")
		{
			// 获取当前用户
			usersGroup.GET("", usersController.Index)
		}

		// 分类
		categoriesController := new(v1controller.CategoriesController)
		categoriesGroup := v1.Group("/categories")
		{
			categoriesGroup.GET("", categoriesController.Index)
			categoriesGroup.POST("", middlewares.AuthJWT(), categoriesController.Store)
			categoriesGroup.PUT("/:id", middlewares.AuthJWT(), categoriesController.Update)
		}

		// 话题
		topicsController := new(v1controller.TopicsController)
		topicsGroup := v1.Group("/topics")
		{
			topicsGroup.POST("", middlewares.AuthJWT(), topicsController.Store)
			topicsGroup.PUT("/:id", middlewares.AuthJWT(), topicsController.Update)
			topicsGroup.DELETE("/:id", middlewares.AuthJWT(), topicsController.Delete)
			topicsGroup.GET("", middlewares.AuthJWT(), topicsController.Index)
			topicsGroup.GET("/:id", middlewares.AuthJWT(), topicsController.Show)
		}

		// 友情连接
		linksController := new(v1controller.LinksController)
		linksGroup := v1.Group("/links")
		{
			linksGroup.GET("", linksController.Index)
		}
	}
}
