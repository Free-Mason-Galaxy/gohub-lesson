// Package database
// descr 数据库操作
// author fm
// date 2022/11/15 10:08
package database

import (
	"database/sql"
	"errors"
	"fmt"

	"gohub-lesson/pkg/config"
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

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMySQLTables()
	case "sqlite":
		err = deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteAllSqliteTables() (err error) {

	var tables []string

	// 读取所有数据表
	err = DB.Select(&tables, `SELECT name FROM sqlite_master WHERE type='table'`).Error

	if err != nil {
		return
	}

	// 删除所有表
	for _, table := range tables {
		err = DB.Migrator().DropTable(table)
		if err != nil {
			return
		}
	}

	return nil
}

func deleteMySQLTables() (err error) {
	var (
		tables []string
		dbname = CurrentDatabase()
	)

	// 读取所有数据表
	err = DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error

	if err != nil {
		return
	}

	// 暂时关闭外键检测
	DB.Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {
		err = DB.Migrator().DropTable(table)
		if err != nil {
			return
		}
	}

	// 开启 MySQL 外键检测
	DB.Exec("SET foreign_key_checks = 1;")

	return nil
}

func TableName(obj any) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(obj)

	return stmt.Schema.Table
}
