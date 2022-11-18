// Package auth
// descr 发送验证码
// author fm
// date 2022/11/17 17:47
package auth

import (
	"github.com/gin-gonic/gin"
	"gohub-lesson/app/requests"
	"gohub-lesson/pkg/captcha"
	"gohub-lesson/pkg/logger"
	baseresponse "gohub-lesson/pkg/response"
	"gohub-lesson/pkg/verifycode"
)

type SendVerifyCodeController struct {
}

// SendEmail 发送邮件
func (class *SendVerifyCodeController) SendEmail(ctx *gin.Context) {

	data, errs := requests.ValidateSendVerifyCodeEmail(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	response := baseresponse.NewResponse(ctx)

	// 发送邮件
	err := verifycode.NewVerifyCode().SendEmail(data.Email)

	if err != nil {
		response.Abort500("发送 Email 验证码失败~")
		return
	}

	response.Success()
}

// ShowCaptcha 获取图片验证码
func (class *SendVerifyCodeController) ShowCaptcha(ctx *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)

	baseresponse.JSON(ctx, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// SendUsingPhone 发送手机验证码
func (class *SendVerifyCodeController) SendUsingPhone(ctx *gin.Context) {

	// 1. 验证表单
	data, errs := requests.ValidateSendVerifyCodePhone(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	response := baseresponse.NewResponse(ctx)

	// 2. 发送 SMS
	if ok := verifycode.NewVerifyCode().SendSMS(data.Phone); !ok {
		response.Abort500("发送短信失败~")
		return
	}

	response.Success()
}
