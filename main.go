// Package gohub_lesson
// descr
// author fm
// date 2022/11/14 15:42
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/arl/statsviz"
	"github.com/spf13/cobra"
	cmd "gohub-lesson/app/cmd"
	cmdMake "gohub-lesson/app/cmd/make"
	"gohub-lesson/bootstrap"
	bsConfig "gohub-lesson/config"
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/console"
)

func main() {

	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = NewRootCmd()

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		cmdMake.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	registerStatsviz()

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

	// 配置初始化，依赖于 --env 参数
	// env := getEnvFlag()

	// config.InitConfig(env)

	// 初始化 gin
	// r := gin.New()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	// gin.SetMode(gin.DebugMode)

	// 初始化 logger
	// bootstrap.SetupLogger()

	// 初始化 redis
	// bootstrap.SetupRedis()

	// 初始化 db
	// bootstrap.SetupDB()

	// 注册路由
	// bootstrap.SetupRoute(r)

	// registerStatsviz()

	// 运行
	// if err := r.Run(config.GetDefaultAddr()); err != nil {
	// 	fmt.Println(err)
	// }
}

// NewRootCmd 创建一个主命令
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "Gohub",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			console.Error("测试 PersistentPreRun")
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
		},
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
