// Package requests
// descr
// author fm
// date 2022/11/15 14:17
package requests

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub-lesson/pkg/captcha"
	"gohub-lesson/pkg/response"
	"gohub-lesson/pkg/verifycode"
)

type BaseRequest struct {
}

// MapErrs 返回数据
type MapErrs struct {
	url.Values
}

var BadRequestErrMsg = "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。"

// Append 追加值
func (class *MapErrs) Append(key, value string) {
	// 使用 append 防止覆盖 key 旧内容
	class.Values[key] = append(class.Values[key], value)
}

// Len 长度
func (class *MapErrs) Len() int {
	return len(class.Values)
}

// IsErrs 是否存在错误
func (class *MapErrs) IsErrs() bool {
	return len(class.Values) > 0
}

// ErrsAbortWithStatusJSON 有错误则 Abort
func (class *MapErrs) ErrsAbortWithStatusJSON(ctx *gin.Context) bool {
	if class.IsErrs() {
		response.ValidationError(ctx, class.Values)
		return true
	}
	return false
}

// ShouldBindQuery 解析数据
// request 引用(指针)
func ShouldBindQuery(request any, ctx *gin.Context) {
	if err := ctx.ShouldBindQuery(request); err != nil {
		response.BadRequest(ctx, err, BadRequestErrMsg)
	}
}

// ShouldBind 解析数据
// request 引用(指针)
func ShouldBind(request any, ctx *gin.Context) {
	if err := ctx.ShouldBind(request); err != nil {
		response.BadRequest(ctx, err, BadRequestErrMsg)
	}
}

// ShouldBindJSON 解析数据
// request 引用(指针)
func ShouldBindJSON(request any, ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(request); err != nil {
		response.BadRequest(ctx, err, BadRequestErrMsg)
	}
}

// Validate 验证
// 用于在控制器层调用
// 如：
//
//	fn(){
//		request := requests.VerifyCodePhoneRequest{}
//		if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
//			return
//		}
//	}
func Validate(ctx *gin.Context, request any, fn func(data any, ctx *gin.Context) map[string][]string) bool {

	ShouldBindJSON(request, ctx)

	errs := fn(request, ctx)

	if len(errs) > 0 {
		response.ValidationError(ctx, errs)
		return false
	}

	return true
}

// validate 内部调用的验证
func validate(data any, rules, messages govalidator.MapData) (errs MapErrs) {

	// 初始化配置
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	errs.Values = govalidator.New(opts).ValidateStruct()

	return
}

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

	ok := verifycode.NewVerifyCode().CheckAnswer(key, answer)

	if !ok {
		errs.Append("verify_code", "验证码错误")
	}

	return errs
}
