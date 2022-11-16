// Package database
// descr 数据库操作
// author fm
// date 2022/11/15 10:08
package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(config gorm.Dialector, logger gormLogger.Interface) {

	var err error
	// 连接数据库
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		fmt.Println(err)
	}

	// 获取底层SQL
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err)
	}
}
