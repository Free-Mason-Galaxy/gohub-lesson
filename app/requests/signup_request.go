// Package requests
// descr 处理请求数据与表单验证
// author fm
// date 2022/11/15 14:14
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	BaseRequest
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExistRequest(ctx *gin.Context) (data SignupPhoneExistRequest, errs MapErrs) {

	ShouldBindJSON(&data, ctx)

	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	// 开始验证并转换
	errs = validate(&data, rules, messages)

	return
}
