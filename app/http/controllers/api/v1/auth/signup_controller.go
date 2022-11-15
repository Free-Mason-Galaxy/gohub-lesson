// Package auth
// descr 处理用户身份认证相关逻辑
// author fm
// date 2022/11/15 11:19
package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "gohub-lesson/app/http/controllers/api/v1"
	"gohub-lesson/app/models/user"
	"gohub-lesson/app/requests"
)

type SignupController struct {
	v1.BaseController
}

func (class *SignupController) IsPhoneExist(ctx *gin.Context) {

	var (
		request = requests.SignupPhoneExistRequest{}
	)

	// 解析 JSON 请求
	// 解析 JSON 请求
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	errs := requests.ValidateSignupPhoneExistRequest(&request, ctx)
	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(request.Phone)})
}
