// Package bootstrap
// descr
// author fm
// date 2022/11/15 17:48
package bootstrap

import (
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/logger"
)

// SetupLogger 初始化日志
func SetupLogger() {

	logger.InitLogger(
		config.GetString("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.GetString("log.type"),
		config.GetString("log.level"),
	)

}
