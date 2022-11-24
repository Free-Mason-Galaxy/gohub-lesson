// Package category 模型
package category

import (
	"gohub-lesson/app/models"
	"gohub-lesson/pkg/database"
)

type Category struct {
	models.BaseModel

	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	models.Timestamps
}

func (category *Category) Create() {
	database.DB.Create(category)
}

func (category *Category) Save() (rowsAffected models.RowsAffected) {
	rowsAffected = models.RowsAffected(database.DB.Save(&category).RowsAffected)
	return
}

func (category *Category) Delete() (rowsAffected models.RowsAffected) {
	rowsAffected = models.RowsAffected(database.DB.Delete(&category).RowsAffected)
	return
}
