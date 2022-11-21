// Package auth
// descr auth 授权相关逻辑
// author fm
// date 2022/11/21 11:17
package auth

import (
	"errors"

	"gohub-lesson/app/models/user"
)

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	userModel := user.GetByMulti(email)

	if userModel.NotExists() {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)

	if userModel.NotExists() {
		return user.User{}, errors.New("手机号未注册")
	}

	return userModel, nil
}
