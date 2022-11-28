// Package link 模型
package link

import (
	"gohub-lesson/app/models"
	"gohub-lesson/pkg/database"
)

type Link struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`

	models.Timestamps
}

func (link *Link) Create() {
	database.DB.Create(&link)
}

func (link *Link) Save() (rowsAffected models.RowsAffected) {
	rowsAffected = models.RowsAffected(database.DB.Save(link).RowsAffected)
	return
}

func (link *Link) Delete() (rowsAffected models.RowsAffected) {
	rowsAffected = models.RowsAffected(database.DB.Delete(link).RowsAffected)
	return
}
