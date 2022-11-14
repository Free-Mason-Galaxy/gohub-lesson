// Package gohub_lesson
// descr
// author fm
// date 2022/11/14 15:42
package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"gohub-lesson/bootstrap"
	bsConfig "gohub-lesson/config"
	"gohub-lesson/pkg/config"
)

func main() {

	// 配置初始化，依赖于 --env 参数
	var env string

	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 配置文件")
	flag.Parse()

	config.InitConfig(env)

	// 初始化 gin
	r := gin.New()

	// 注册路由
	bootstrap.SetupRoute(r)

	// 运行
	if err := r.Run(config.GetDefaultAddr()); err != nil {
		fmt.Println(err)
	}
}

func init() {
	bsConfig.Initialize()
}
