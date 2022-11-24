// Package category 工具
package category

import (
    "gohub-lesson/pkg/database"
    "gohub-lesson/pkg/paginator"
    "gohub-lesson/pkg/app"

    "gorm.io/gorm/clause"
    "github.com/gin-gonic/gin"
)

func Get(idStr string) (category Category) {
    database.DB.Where("id", idStr).First(&category)
    return
}

func GetBy(field, value string) (category Category) {
    database.DB.Where("? = ?", clause.Column{Name: field}, value).First(&category)
    return
}

func All() (categories []Category) {
    database.DB.Find(&categories)
    return
}

func IsExist(field, value string) bool {
    var m Category

    database.DB.Select("id").
        Where("? = ?", clause.Column{Name: field}, value).
        Take(&m)

    return m.ID > 0
}

func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Category{}),
        &categories,
        app.V1URL(database.TableName(&Category{})),
        perPage,
    )
    return
}