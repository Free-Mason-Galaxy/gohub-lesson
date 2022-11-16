// Package test
// descr
// author fm
// date 2022/11/16 10:45
package test

import (
	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func Test() {
	panic("这是panic测试代码")
}

func (class *TestController) Any(ctx *gin.Context) {
	Test()
}
