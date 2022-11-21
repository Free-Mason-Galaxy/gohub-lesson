// Package auth
// descr
// author fm
// date 2022/11/21 11:26
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub-lesson/app/http/controllers/api/v1"
	"gohub-lesson/app/requests"
	"gohub-lesson/pkg/auth"
	"gohub-lesson/pkg/helpers"
	"gohub-lesson/pkg/jwt"
	"gohub-lesson/pkg/response"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(ctx *gin.Context) {

	// 1. 验证表单
	data, errs := requests.ValidateLoginByPhone(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	// 2. 尝试登录
	user, err := auth.LoginByPhone(data.Phone)

	if helpers.IsError(err) {
		// 失败，显示错误提示
		response.Error(ctx, err, "账号不存在")
		return
	}

	// 登录成功
	token := jwt.NewJWT().GenerateToken(user.GetIdString(), user.Name)

	response.JSON(ctx, gin.H{
		"token": token,
	})
}

// LoginByPassword 多种方法登录，支持手机号、email 和用户名
func (lc *LoginController) LoginByPassword(ctx *gin.Context) {
	// 1. 验证表单
	data, errs := requests.ValidateLoginByPassword(ctx)

	if errs.ErrsAbortWithStatusJSON(ctx) {
		return
	}

	// 2. 尝试登录
	user, err := auth.Attempt(data.LoginID, data.Password)

	if helpers.IsError(err) {
		// 失败，显示错误提示
		response.Unauthorized(ctx, "账号不存在或密码错误")
		return
	}

	token := jwt.NewJWT().GenerateToken(user.GetIdString(), user.Name)

	response.JSON(ctx, gin.H{
		"token": token,
	})
}

// RefreshToken 重置 token
func (class *LoginController) RefreshToken(ctx *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(ctx)

	if helpers.IsError(err) {
		response.Error(ctx, err, "令牌刷新失败")
		return
	}

	response.JSON(ctx, gin.H{
		"token": token,
	})
}
