package migrations

import (
	"database/sql"

	"gohub-lesson/app/models"
	"gohub-lesson/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Categories struct {
		models.BaseModel
		Name        string `gorm:"type:varchar(255);not null;index"`
		Description string `gorm:"type:varchar(255);default:null"`
		models.Timestamps
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Categories{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Categories{})
	}

	migrate.Add("2022_11_24_152056_add_categories_table", up, down)
}
