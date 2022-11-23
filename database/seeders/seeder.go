// Package seeders
// descr
// author fm
// date 2022/11/23 17:04
package seeders

import (
	"gohub-lesson/pkg/seed"
)

// Initialize 初始化
func Initialize() {

	// 触发加载本目录下其他文件中的 init 方法

	// 指定优先于同目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
