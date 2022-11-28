// Package seeders
package seeders

import (
	"fmt"

	"gohub-lesson/database/factories"
	"gohub-lesson/pkg/console"
	"gohub-lesson/pkg/logger"
	"gohub-lesson/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	// 添加 Seeder
	seed.Add("SeedTopicsTable", func(db *gorm.DB) {

		// 创建 10 个用户对象
		Topics := factories.MakeTopics(10)

		// 批量创建用户（注意批量创建不会调用模型钩子）
		result := db.Table("topics").Create(&Topics)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
