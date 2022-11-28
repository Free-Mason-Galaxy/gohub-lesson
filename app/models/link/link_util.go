// Package link 工具
package link

import (
	"time"

	"gohub-lesson/pkg/app"
	"gohub-lesson/pkg/cache"
	"gohub-lesson/pkg/database"
	"gohub-lesson/pkg/helpers"
	"gohub-lesson/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func AllCached() (links []Link) {
	// 设置缓存 key
	cacheKey := "links:all"

	// 设置过期时间
	expireTime := 120 * time.Minute

	// 取数据
	cache.GetObject(cacheKey, &links)

	// 如果数据为空
	if helpers.Empty(links) {
		// 查询数据库
		links = All()
		if helpers.Empty(links) {
			return links
		}
		// 设置缓存
		cache.Set(cacheKey, links, expireTime)
	}

	return
}

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
