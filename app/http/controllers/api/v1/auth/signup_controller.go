// Package auth
// descr 处理用户身份认证相关逻辑
// author fm
// date 2022/11/15 11:19
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub-lesson/app/http/controllers/api/v1"
	userModel "gohub-lesson/app/models/user"
	"gohub-lesson/app/requests"
	baseresponse "gohub-lesson/pkg/response"
)

type SignupController struct {
	v1.BaseController
}

// SignupUsingEmail 使用手机号注册
func (class *SignupController) SignupUsingEmail(ctx *gin.Context) {

	data, errs := requests.ValidateSignupUsingEmail(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	user := userModel.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	user.Create()

	response := baseresponse.NewResponse(ctx)

	if user.Exists() {
		response.CreatedJSON(gin.H{
			"data": user,
		})
		return
	}

	response.Abort500("创建用户失败，请稍后再试~")

	return
}

// SignupUsingPhone 使用手机号注册
func (class *SignupController) SignupUsingPhone(ctx *gin.Context) {

	data, errs := requests.ValidateSignupUsingPhone(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	user := userModel.User{
		Name:     data.Name,
		Phone:    data.Phone,
		Password: data.Password,
	}

	user.Create()

	response := baseresponse.NewResponse(ctx)

	if user.Exists() {
		response.CreatedJSON(gin.H{
			"data": user,
		})
		return
	}

	response.Abort500("创建用户失败，请稍后再试~")

	return
}

// IsPhoneExist 判断手机号是否存在
func (class *SignupController) IsPhoneExist(ctx *gin.Context) {

	data, errs := requests.ValidateSignupPhoneExist(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	baseresponse.JSON(ctx, gin.H{"exist": userModel.IsPhoneExist(data.Phone)})
}

// IsEmailExist 判断邮箱是否存在
func (class *SignupController) IsEmailExist(ctx *gin.Context) {

	data, errs := requests.ValidateSignupEmailExist(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	baseresponse.JSON(ctx, gin.H{"exist": userModel.IsEmailExist(data.Email)})
}
