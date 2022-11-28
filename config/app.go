// Package config
// descr 项目配置信息
// author fm
// date 2022/11/14 17:09
package config

import (
	"gohub-lesson/pkg/config"
)

func init() {
	config.Add("app", func() map[string]any {
		return map[string]any{
			// 应用名称
			"name": config.Env("APP_NAME", "gohub"),

			// 当前环境，用以区分多环境，一般为 local, stage, production, test
			"env": config.Env("APP_ENV", "production"),

			// 是否进入调试模式
			"debug": config.Env("APP_DEBUG", false),

			// 应用服务端口
			"port": config.Env("APP_PORT", "82"),

			// 加密会话、JWT 加密
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

			// 用以生成链接
			"url": config.Env("APP_URL", "http://localhost:82"),

			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),

			// API 域名，api 如 http://api.domain.com/v1/users
			// 未设置的话所有 API URL 加 api 前缀，如 http://domain.com/api/v1/users
			"api_domain": config.Env("API_DOMAIN"),
		}
	})
}
