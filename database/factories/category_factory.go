// Package factories
package factories

import (
	"gohub-lesson/app/models/category"

	"github.com/bxcodec/faker/v4"
)

func MakeCategories(times int) []category.Category {

	var objs []category.Category

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := category.Category{
			Name:        faker.Username(),
			Description: faker.Sentence(),
		}
		objs = append(objs, model)
	}

	return objs
}
