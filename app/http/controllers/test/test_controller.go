// Package test
// descr
// author fm
// date 2022/11/16 10:45
package test

import (
	"github.com/gin-gonic/gin"
	"gohub-lesson/pkg/captcha"
	"gohub-lesson/pkg/logger"
	"gohub-lesson/pkg/response"
)

type TestController struct {
}

func Test() {
	panic("这是panic测试代码")
}

func (class *TestController) Any(ctx *gin.Context) {
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
