// Package factories
// descr 存放工厂方法
// author fm
// date 2022/11/23 16:58
package factories

import (
	"gohub-lesson/app/models/user"
	"gohub-lesson/pkg/helpers"

	"github.com/bxcodec/faker/v4"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$yNqswpIx3PLywaI5f71qgeH0ItKEs.7x8e1IUFsTLwcHUB4ZVeen.",
		}
		objs = append(objs, model)
	}

	return objs
}
