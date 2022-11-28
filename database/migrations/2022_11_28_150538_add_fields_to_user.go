package migrations

import (
	"database/sql"

	"gohub-lesson/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		City         string `json:"city,omitempty"`
		Introduction string `json:"introduction,omitempty"`
		Avatar       string `json:"avatar,omitempty"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&User{}, "City")
		migrator.DropColumn(&User{}, "Introduction")
		migrator.DropColumn(&User{}, "Avatar")
	}

	migrate.Add("2022_11_28_150538_add_fields_to_user", up, down)
}
