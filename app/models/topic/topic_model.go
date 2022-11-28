// Package topic 模型
package topic

import (
	"gohub-lesson/app/models"
	"gohub-lesson/app/models/category"
	"gohub-lesson/app/models/user"
	"gohub-lesson/pkg/database"
)

type Topic struct {
	// 字段
	models.BaseModel

	Title      string `json:"title,omitempty" `
	Body       string `json:"body,omitempty" `
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	models.Timestamps

	// 关联数据
	User     user.User         `json:"user"`
	Category category.Category `json:"category"`
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected models.RowsAffected) {
	rowsAffected = models.RowsAffected(database.DB.Save(topic).RowsAffected)
	return
}

func (topic *Topic) Delete() (rowsAffected models.RowsAffected) {
	rowsAffected = models.RowsAffected(database.DB.Delete(topic).RowsAffected)
	return
}
