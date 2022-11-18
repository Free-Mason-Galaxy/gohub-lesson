// Package requests
// descr
// author fm
// date 2022/11/17 17:25
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub-lesson/pkg/captcha"
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

	if !captcha.NewCaptcha().VerifyCaptcha(data.CaptchaID, data.CaptchaAnswer) {
		errs.Append("captcha_answer", "图片验证码错误")
	}

	return
}
