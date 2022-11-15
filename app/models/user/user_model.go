// Package user
// descr 用户模型
// author fm
// date 2022/11/15 10:51
package user

import (
	"gohub-lesson/app/models"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty" gorm:"type:varchar(191)"`
	Email    string `json:"-" gorm:"type:varchar(191)"`
	Phone    string `json:"-" gorm:"type:varchar(191)"`
	Password string `json:"-" gorm:"type:varchar(191)"`

	models.Timestamps
}
