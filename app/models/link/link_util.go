// Package link 工具
package link

import (
    "gohub-lesson/pkg/database"
    "gohub-lesson/pkg/paginator"
    "gohub-lesson/pkg/app"

    "gorm.io/gorm/clause"
    "github.com/gin-gonic/gin"
)

func Get(idStr string) (link Link) {
    database.DB.Where("id", idStr).First(&link)
    return
}

func GetBy(field, value string) (link Link) {
    database.DB.Where("? = ?", clause.Column{Name: field}, value).First(&link)
    return
}

func All() (links []Link) {
    database.DB.Find(&links)
    return
}

func IsExist(field, value string) bool {
    var m Link

    database.DB.Select("id").
        Where("? = ?", clause.Column{Name: field}, value).
        Take(&m)

    return m.ID > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Link{}),
        &links,
        app.V1URL(database.TableName(&Link{})),
        perPage,
    )
    return
}