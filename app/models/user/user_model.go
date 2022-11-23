// Package user
// descr 用户模型
// author fm
// date 2022/11/15 10:51
package user

import (
	"gohub-lesson/app/models"
	"gohub-lesson/pkg/database"
	"gohub-lesson/pkg/hash"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty" gorm:"type:varchar(191)"`
	Email    string `json:"-" gorm:"type:varchar(191)"`
	Phone    string `json:"-" gorm:"type:varchar(191)"`
	Password string `json:"-" gorm:"type:varchar(191)"`

	models.Timestamps
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (class *User) Create() {
	database.DB.Create(class)
}

// ComparePassword 密码是否正确
func (class *User) ComparePassword(pwd string) bool {
	return hash.BcryptCheck(pwd, class.Password)
}

// Save 更新
func (class *User) Save() int64 {
	return database.DB.Save(class).RowsAffected
}

// Delete 删除
func (class *User) Delete() (rowsAffected int64) {
	return database.DB.Delete(class).RowsAffected
}
