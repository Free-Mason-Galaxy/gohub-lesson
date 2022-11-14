// Package bootstrap
// descr 处理程序初始化逻辑
// author fm
// date 2022/11/14 16:14
package bootstrap

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gohub-lesson/routes"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleware(router)

	// 注册 API 路由
	routes.RegisterAPIRoutes(router)

	// 配置 404
	setup404Handler(router)
}

// setup404Handler 处理 404
func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		accept := ctx.GetHeader("Accept")
		if strings.Contains(accept, "text/html") {
			ctx.String(http.StatusNotFound, "Not Found 404")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_msg":  "路由未定义",
			})
		}
	})
}

// registerGlobalMiddleware 注册全局中间件
func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}
