// Package gohub_lesson
// descr
// author fm
// date 2022/11/14 15:42
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gohub-lesson/bootstrap"
)

func main() {

	// 初始化 gin
	r := gin.Default()

	// 注册路由
	bootstrap.SetupRoute(r)

	// 运行
	if err := r.Run(":82"); err != nil {
		fmt.Println(err)
	}
}
