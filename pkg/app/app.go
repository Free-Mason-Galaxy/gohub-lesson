// Package app
// descr 应用信息
// author fm
// date 2022/11/15 17:10
package app

import (
	"gohub-lesson/pkg/config"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}
