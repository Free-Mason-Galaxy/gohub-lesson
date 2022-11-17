// Package auth
// descr
// author fm
// date 2022/11/16 17:56
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub-lesson/app/http/controllers/api/v1"
	"gohub-lesson/pkg/captcha"
	"gohub-lesson/pkg/logger"
	"gohub-lesson/pkg/response"
)

type VerifyCodeController struct {
	v1.BaseController
}

func (class *VerifyCodeController) ShowCaptcha(ctx *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)

	response.JSON(ctx, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
