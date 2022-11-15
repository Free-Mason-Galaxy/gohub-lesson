// Package auth
// descr 处理用户身份认证相关逻辑
// author fm
// date 2022/11/15 11:19
package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "gohub-lesson/app/http/controllers/api/v1"
	"gohub-lesson/app/models/user"
	"gohub-lesson/app/requests"
)

type SignupController struct {
	v1.BaseController
}

// IsPhoneExist 判断手机号是否存在
func (class *SignupController) IsPhoneExist(ctx *gin.Context) {

	data, errs := requests.ValidateSignupPhoneExistRequest(ctx)

	if errs.IsErrs() {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(data.Phone)})
}

func (class *SignupController) IsEmailExist(ctx *gin.Context) {

	data, errs := requests.ValidateSignupEmailExist(ctx)

	errs.ErrsAbortWithStatusJSON(ctx)

	ctx.JSON(http.StatusOK, gin.H{"exist": user.IsEmailExist(data.Email)})
}
