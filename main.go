// Package gohub_lesson
// descr
// author fm
// date 2022/11/14 15:42
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
	"gohub-lesson/bootstrap"
	bsConfig "gohub-lesson/config"
	"gohub-lesson/pkg/config"
)

func main() {

	// 配置初始化，依赖于 --env 参数
	env := getEnvFlag()

	config.InitConfig(env)

	// 初始化 gin
	r := gin.New()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.DebugMode)

	// 初始化 logger
	bootstrap.SetupLogger()

	// 初始化 redis
	bootstrap.SetupRedis()

	// 初始化 db
	bootstrap.SetupDB()

	// 注册路由
	bootstrap.SetupRoute(r)

	registerStatsviz()

	// 运行
	if err := r.Run(config.GetDefaultAddr()); err != nil {
		fmt.Println(err)
	}
}

// registerStatsviz 实时可视化Go Runtime指标
func registerStatsviz() {
	statsviz.RegisterDefault()
	go http.ListenAndServe("localhost:62", nil)
}

func getEnvFlag() (env string) {
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 配置文件")
	flag.Parse()
	return
}

func init() {
	bsConfig.Initialize()
}
