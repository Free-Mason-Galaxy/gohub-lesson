// Package bootstrap
// descr 初始化函数
// author fm
// date 2022/11/16 16:26
package bootstrap

import (
	"fmt"

	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/redis"
)

// SetupRedis 初始化 Redis
func SetupRedis() {
	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
