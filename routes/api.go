// Package routes
// descr 注册路由
// author fm
// date 2022/11/14 16:22
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1controller "gohub-lesson/app/http/controllers/api/v1"
	"gohub-lesson/app/http/controllers/api/v1/auth"
	"gohub-lesson/app/http/controllers/test"
	"gohub-lesson/app/http/middlewares"
	pkgAuth "gohub-lesson/pkg/auth"
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/response"
)

func RegisterAPIRoutes(r *gin.Engine) {

	// 静态图片访问
	r.StaticFS("public/uploads", http.Dir("./public/uploads"))
	// r.Static("/assets", "./assets")
	// r.StaticFS("/more_static", http.Dir("my_file_system"))
	// r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.GET("/test_auth", middlewares.AuthJWT(), func(ctx *gin.Context) {
		userModel := pkgAuth.CurrentUser(ctx)
		response.Data(ctx, userModel)
	})

	r.GET("/test_guest", middlewares.GuestJWT(), func(ctx *gin.Context) {
		// userModel := pkgAuth.CurrentUser(ctx)
		response.Success(ctx)
	})

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}
	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。

	t := v1.Group("/")
	{
		testController := new(test.TestController)
		t.Any("/test", testController.Any)
	}

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
			// 更新当前用户
			usersGroup.PUT("", middlewares.AuthJWT(), usersController.UpdateProfile)
			usersGroup.PUT("/email", middlewares.AuthJWT(), usersController.UpdateEmail)
			usersGroup.PUT("/phone", middlewares.AuthJWT(), usersController.UpdatePhone)
			usersGroup.PUT("/password", middlewares.AuthJWT(), usersController.UpdatePassword)
			usersGroup.PUT("/avatar", middlewares.AuthJWT(), usersController.UpdateAvatar)
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
