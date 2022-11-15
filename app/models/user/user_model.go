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

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.Timestamps
}
