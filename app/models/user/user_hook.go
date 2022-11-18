// Package user
// descr 钩子
// author fm
// date 2022/11/18 17:21
package user

import (
	"gohub-lesson/pkg/hash"
	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (class *User) BeforeSave(tx *gorm.DB) (err error) {
	if hash.BcryptIsHashed(class.Password) {
		return
	}
	class.Password = hash.BcryptHash(class.Password)
	return
}
