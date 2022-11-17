// Package test
// descr
// author fm
// date 2022/11/16 10:45
package test

import (
	"github.com/gin-gonic/gin"
	"gohub-lesson/pkg/captcha"
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/logger"
	"gohub-lesson/pkg/response"
	"gohub-lesson/pkg/sms"
)

type TestController struct {
}

func Test() {
	panic("这是panic测试代码")
}

func (class *TestController) Any(ctx *gin.Context) {
	{
		sms.NewSMS().Send("17602118840", sms.Message{
			Template: config.GetString("sms.aliyun.template_code"),
			Data:     map[string]string{"code": "23456"},
		})
		return
	}
	{
		response.CJSON(ctx, 200, nil)
		return
	}
	{
		logger.Dump(captcha.NewCaptcha().VerifyCaptcha(ctx.Query("key"), ctx.Query("value")))
		return
	}
	Test()
}
