// Package test
// descr
// author fm
// date 2022/11/16 10:45
package test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gohub-lesson/app/models/user"
	"gohub-lesson/pkg/captcha"
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/logger"
	"gohub-lesson/pkg/response"
	"gohub-lesson/pkg/sms"
	"gohub-lesson/pkg/verifycode"
)

type TestController struct {
}

func Test() {
	panic("这是panic测试代码")
}

type TestInterface interface {
	Testing() string
}

type T struct {
}

func (class T) Testing() string {
	return "testing"
}

type T2 struct {
	T
}

func fn(t TestInterface) {
	t.Testing()
}

func (class *TestController) Any(ctx *gin.Context) {
	{
		var a = map[string]string{
			"k1": "v1",
			"k2": "v2",
		}
		for k, v := range a {
			fmt.Println("k_addr:", &k, k)
			fmt.Println("v_addr:", &v, v)
		}

		return

	}
	{
		var u user.User
		u.Name = "name1"
		u.Email = "name@qq.com"
		u.Password = "admin"
		u.Create()

		response.JSON(ctx, gin.H{
			"data": u,
		})
		return
	}
	{
		var t T2
		fn(t)
	}
	{
		isSuccess := verifycode.NewVerifyCode().SendSMS(ctx.Query("key"))
		response.JSON(ctx, gin.H{"isSuccess": isSuccess})
		return
	}
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
