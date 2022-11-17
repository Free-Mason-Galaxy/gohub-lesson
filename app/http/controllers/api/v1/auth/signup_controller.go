// Package auth
// descr 处理用户身份认证相关逻辑
// author fm
// date 2022/11/15 11:19
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub-lesson/app/http/controllers/api/v1"
	"gohub-lesson/app/models/user"
	"gohub-lesson/app/requests"
	"gohub-lesson/pkg/response"
)

type SignupController struct {
	v1.BaseController
}

// IsPhoneExist 判断手机号是否存在
func (class *SignupController) IsPhoneExist(ctx *gin.Context) {

	data, errs := requests.ValidateSignupPhoneExistRequest(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	response.JSON(ctx, gin.H{"exist": user.IsPhoneExist(data.Phone)})
}

// IsEmailExist 判断邮箱是否存在
func (class *SignupController) IsEmailExist(ctx *gin.Context) {

	data, errs := requests.ValidateSignupEmailExist(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	response.JSON(ctx, gin.H{"exist": user.IsEmailExist(data.Email)})
}
