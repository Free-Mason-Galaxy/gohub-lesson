// Package test 模型
package test

import (
	"gohub-lesson/app/models"
	"gohub-lesson/pkg/database"
	"gorm.io/gorm"
)

type Test struct {
	models.BaseModel

	// 模型模板中放进去常用的方法，使用 FIXME() 这个不存在的函数，通知要记得修改这个地方
	Title string `json:"title" gorm:"column:title"`
	models.Timestamps

	// Transaction 事务
	Transaction *gorm.DB `json:"-" gorm:"-"`
}

func (test *Test) Create() {
	database.DB.Create(&test)
}

func (test *Test) Save() (rowsAffected int64) {
	result := database.DB.Save(&test)
	return result.RowsAffected
}

func (test *Test) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&test)
	return result.RowsAffected
}

func (test *Test) TxCreate() *gorm.DB {
	return test.Transaction.Create(test)
}

func (test *Test) TxSave() (rowsAffected int64) {
	result := test.Transaction.Save(test)
	return result.RowsAffected
}

func (test *Test) TxDelete() (rowsAffected int64) {
	result := test.Transaction.Delete(test)
	return result.RowsAffected
}
