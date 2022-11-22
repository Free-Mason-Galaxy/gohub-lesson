// Package cmd
// descr
// 方便临时调试代码
// 之前都在 main.go 里调试，有时候代码会忘记删除，在 play 命令里测试代码，忘记删除了也不用担心影响到主程序。
// Play 命令有点像 go.dev/play/ ，但是运行在应用环境中，数据库、配置、Redis 等系统服务都已初始化，可以放心使用。
//
// author fm
// date 2022/11/22 14:42
package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"gohub-lesson/pkg/console"
	"gohub-lesson/pkg/redis"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// runPlay 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	// 存进 redis 中
	redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	// 从 redis 里取出
	console.Success(redis.Redis.Get("hello"))
	console.Error("测试 console 输出颜色")
}
