// Package requests
// descr
// author fm
// date 2022/11/17 17:25
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SendVerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSendVerifyCodePhone(ctx *gin.Context) (data SendVerifyCodePhoneRequest, errs MapErrs) {

	ShouldBindJSON(&data, ctx)

	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号必填，参数名称 phone",
			"digits:手机号长度必须 11 位的数字",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 必填",
		},
		"captcha_answer": []string{
			"required:图片验证码必须",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}

	errs = validate(&data, rules, messages)

	ValidateCaptcha(data.CaptchaID, data.CaptchaAnswer, errs)

	return
}

type SendVerifyCodeEmailRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Email string `json:"email,omitempty" valid:"email"`
}

// ValidateSendVerifyCodeEmail 验证表单，返回长度等于零即通过
func ValidateSendVerifyCodeEmail(ctx *gin.Context) (data SendVerifyCodeEmailRequest, errs MapErrs) {

	ShouldBindJSON(&data, ctx)

	// 1. 定制认证规则
	rules := govalidator.MapData{
		"email":          []string{"required", "min:4", "max:30", "email"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}

	errs = validate(&data, rules, messages)

	// 图片验证码
	errs = ValidateCaptcha(data.CaptchaID, data.CaptchaAnswer, errs)

	return
}
