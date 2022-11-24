package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CategoryRequest struct {
	Name  string `valid:"name" json:"name"`
	Descr string `valid:"descr" json:"descr,omitempty"`
}

// ValidateCategory 验证表单，返回长度等于零即通过
func ValidateCategory(ctx *gin.Context) (data CategoryRequest, errs MapErrs) {

	ShouldBindJSON(&data, ctx)

	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
		"description": []string{"min_cn:3", "max_cn:255"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min_cn:分类名称长度需至少 2 个字",
			"max_cn:分类名称长度不能超过 8 个字",
			"not_exists:分类名称已存在",
		},
		"description": []string{
			"min_cn:分类描述长度需至少 3 个字",
			"max_cn:分类描述长度不能超过 255 个字",
		},
	}

	errs = validate(&data, rules, messages)

	return
}
