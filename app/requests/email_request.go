// Package requests
// descr
// author fm
// date 2022/11/15 15:09
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type EmailRequest struct {
	BaseRequest
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateSignupEmailExist(ctx *gin.Context) (data EmailRequest, errs MapErrs) {
	// 获取提交数据
	// 把提交的数据解析到 toData
	ShouldBindJSON(&data, ctx)

	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义错误
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 必须填写",
			"min:Email 最少 4 个字符",
			"max:Email 最多 30 个字符",
			"email:Email 格式不正确，请输入有效的邮箱地址",
		},
	}

	// 开始验证并转换
	errs = validate(&data, rules, messages)

	return
}
