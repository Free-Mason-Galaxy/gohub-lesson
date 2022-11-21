// Package auth
// descr
// author fm
// date 2022/11/21 15:15
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub-lesson/app/http/controllers/api/v1"
	"gohub-lesson/app/models/user"
	"gohub-lesson/app/requests"
	"gohub-lesson/pkg/response"
)

type PasswordController struct {
	v1.BaseController
}

// ResetByPhone 使用手机重置密码
func (class *PasswordController) ResetByPhone(ctx *gin.Context) {
	// 1. 验证表单
	data, errs := requests.ValidateResetByPhone(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	// 2. 更新密码
	userModel := user.GetByPhone(data.Phone)

	if userModel.NotExists() {
		response.Abort404(ctx)
		return
	}

	userModel.Password = data.Password
	userModel.Save()

	response.Success(ctx)
}
