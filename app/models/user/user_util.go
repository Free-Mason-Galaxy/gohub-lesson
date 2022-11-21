// Package user
// descr 模型相关的数据库操作
// author fm
// date 2022/11/15 11:04
package user

import (
	"gohub-lesson/pkg/database"
	"gorm.io/gorm/clause"
)

// IsEmailExist 判断 email 是否被注册(是否存在)
func IsEmailExist(email string) bool {
	return IsColumnValueExist("`email`", email)
}

// IsPhoneExist 判断 phone 是否被注册(是否存在)
func IsPhoneExist(phone string) bool {
	return IsColumnValueExist("phone", phone)
}

// IsColumnValueExist 判断字段值是否存在
func IsColumnValueExist(column string, value string) bool {
	var m User

	database.DB.Model(User{}).
		Select("id").
		Where("? = ?", clause.Column{Name: column}, value).
		Take(&m)

	return m.ID > 0
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel User) {
	database.DB.
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&userModel)
	return
}
