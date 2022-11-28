// Package factories
package factories

import (
	"gohub-lesson/app/models/link"

	"github.com/bxcodec/faker/v4"
)

func MakeLinks(times int) []link.Link {

	var objs []link.Link

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := link.Link{
			Name: faker.Username(),
			URL:  faker.URL(),
		}
		objs = append(objs, model)
	}

	return objs
}
