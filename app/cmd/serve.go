// Package cmd
// descr
// author fm
// date 2022/11/22 13:48
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gohub-lesson/bootstrap"
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/console"
	"gohub-lesson/pkg/logger"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,

	// Args Cobra 提供了一组参数校验器，用以快速校验命令的参数是否符合预期，否则报错：
	//
	// NoArgs - 如果存在任何位置参数，该命令将报错
	// ArbitraryArgs - 该命令会接受任何位置参数
	// OnlyValidArgs - 如果有任何位置参数不在命令的 ValidArgs 字段中，该命令将报错
	// MinimumNArgs(int) - 至少要有 N 个位置参数，否则报错
	// MaximumNArgs(int) - 如果位置参数超过 N 个将报错
	// ExactArgs(int) - 必须有 N 个位置参数，否则报错
	// `ExactValidArgs (int) 必须有 N 个位置参数，且都在命令的 ValidArgs 字段中，否则报错
	// RangeArgs(min, max) - 如果位置参数的个数不在区间 min 和 max 之中，报错
	// MatchAll(pargs ...PositionalArgs) - 支持使用以上的多个验证器
	Args: cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(r)

	// 运行服务器
	err := r.Run(":" + config.Get("app.port"))

	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}
