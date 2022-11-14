// Package routes
// descr 注册路由
// author fm
// date 2022/11/14 16:22
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"test": "hello world",
			})
		})
	}
}
