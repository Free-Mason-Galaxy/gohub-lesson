// Package test 工具
package test

import (
    "gohub-lesson/pkg/database"
    "gohub-lesson/pkg/paginator"
    "gohub-lesson/pkg/app"

    "gorm.io/gorm/clause"
    "github.com/gin-gonic/gin"
)

func Get(idStr string) (test Test) {
    database.DB.Where("id", idStr).First(&test)
    return
}

func GetBy(field, value string) (test Test) {
    database.DB.Where("? = ?", clause.Column{Name: field}, value).First(&test)
    return
}

func All() (tests []Test) {
    database.DB.Find(&tests)
    return
}

func IsExist(field, value string) bool {
    var m Test

    database.DB.Select("id").
        Where("? = ?", clause.Column{Name: field}, value).
        Take(&m)

    return m.ID > 0
}

func Paginate(c *gin.Context, perPage int) (tests []Test, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Test{}),
        &tests,
        app.V1URL(database.TableName(&Test{})),
        perPage,
    )
    return
}