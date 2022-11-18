// Package requests
// descr 自定义验证规则
// author fm
// date 2022/11/18 14:57
package requests

import (
	"gohub-lesson/pkg/captcha"
	"gohub-lesson/pkg/verifycode"
)

// ValidateCaptcha 自定义规则，验证『图片验证码』
func ValidateCaptcha(captchaID, captchaAnswer string, errs MapErrs) MapErrs {
	// if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
	// 	errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	// }
	if !captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer) {
		errs.Append("captcha_answer", "图片验证码错误")
	}
	return errs
}

// ValidatePasswordConfirm 自定义规则，检查两次密码是否正确
func ValidatePasswordConfirm(password, passwordConfirm string, errs MapErrs) MapErrs {
	if password != passwordConfirm {
		errs.Append("password_confirm", "两次输入密码不匹配!")
	}
	return errs
}

// ValidateVerifyCode 自定义规则，验证『手机/邮箱验证码』
func ValidateVerifyCode(key, answer string, errs MapErrs) MapErrs {
	if ok := verifycode.NewVerifyCode().CheckAnswer(key, answer); !ok {
		errs.Append("verify_code", "验证码错误")
	}
	return errs
}
