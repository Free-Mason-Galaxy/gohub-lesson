// Package requests
// descr
// author fm
// date 2022/11/15 14:17
package requests

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type BaseRequest struct {
}

// MapErrs 返回数据
type MapErrs struct {
	url.Values
}

// Len 长度
func (class *MapErrs) Len() int {
	return len(class.Values)
}

// IsErrs 是否存在错误
func (class *MapErrs) IsErrs() bool {
	return len(class.Values) > 0
}

// ErrsAbortWithStatusJSON 有错误则 Abort
func (class *MapErrs) ErrsAbortWithStatusJSON(ctx *gin.Context) bool {
	if class.IsErrs() {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": class.Values,
		})
		return true
	}
	return false
}

// ShouldBindJSON 解析数据
// request 引用(指针)
func ShouldBindJSON(request any, ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}
}

// validate 验证
func validate(data any, rules, messages govalidator.MapData) (errs MapErrs) {

	// 初始化配置
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	errs.Values = govalidator.New(opts).ValidateStruct()

	return
}
