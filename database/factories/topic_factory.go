// Package factories
package factories

import (
	"gohub-lesson/app/models/topic"

	"github.com/bxcodec/faker/v4"
)

func MakeTopics(times int) []topic.Topic {

	var objs []topic.Topic

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := topic.Topic{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			CategoryID: "3",
			UserID:     "1",
		}
		objs = append(objs, model)
	}

	return objs
}
