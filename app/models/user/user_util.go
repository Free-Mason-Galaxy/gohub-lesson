// Package user
// descr 模型相关的数据库操作
// author fm
// date 2022/11/15 11:04
package user

import (
	"gohub-lesson/pkg/database"
)

// IsEmailExist 判断 email 是否被注册(是否存在)
func IsEmailExist(email string) bool {
	return IsColumnExist("email", email)
}

// IsPhoneExist 判断 phone 是否被注册(是否存在)
func IsPhoneExist(phone string) bool {
	return IsColumnExist("phone", phone)
}

// IsColumnExist 判断指定字段值是否存在
func IsColumnExist(column string, value string) bool {
	var m User
	database.DB.Model(User{}).Select("id").Where("? = ?", column, value).First(&m)
	return m.ID > 0
}
