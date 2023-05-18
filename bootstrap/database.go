// Package bootstrap
// descr
// author fm
// date 2022/11/15 10:16
package bootstrap

import (
	"errors"
	"fmt"
	"time"

	"gohub-lesson/app/models/user"
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/database"
	"gohub-lesson/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupDB 初始化数据库
func SetupDB() {

	var dbConfig gorm.Dialector

	switch config.Get("database.connection") {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite
		dbase := config.Get("database.sqlite.database")
		dbConfig = sqlite.Open(dbase)
	default:
		panic(errors.New("database connection not supported"))

	}

	database.Connect(dbConfig, logger.NewGormLogger())

	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))

	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))

	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	// AutoMigrate()

}

func AutoMigrate() {
	database.DB.AutoMigrate(&user.User{})
}
