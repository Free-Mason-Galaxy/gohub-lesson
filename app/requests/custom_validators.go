// Package validators
// descr 自定义验证规则
// author fm
// date 2022/11/18 14:57
package requests

import (
	"gohub-lesson/pkg/captcha"
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
