// Package config
// descr email 配置信息
// author fm
// date 2022/11/18 14:36
package config

import (
	"gohub-lesson/pkg/config"
)

func init() {
	config.Add("mail", func() map[string]any {
		return map[string]interface{}{

			// 默认是 MailHog 的配置
			"smtp": map[string]interface{}{
				"host":     config.Env("MAIL_HOST", "localhost"),
				"port":     config.Env("MAIL_PORT", 1025),
				"username": config.Env("MAIL_USERNAME", ""),
				"password": config.Env("MAIL_PASSWORD", ""),
			},

			"from": map[string]interface{}{
				"address": config.Env("MAIL_FROM_ADDRESS", "gohub-lesson@example.com"),
				"name":    config.Env("MAIL_FROM_NAME", "gohub-lesson"),
			},
		}
	})
}
