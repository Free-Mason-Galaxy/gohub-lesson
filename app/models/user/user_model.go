// Package user
// descr 用户模型
// author fm
// date 2022/11/15 10:51
package user

import (
	"gohub-lesson/app/models"
	"gohub-lesson/pkg/database"
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
