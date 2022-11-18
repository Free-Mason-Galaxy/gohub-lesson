// Package requests
// descr
// author fm
// date 2022/11/15 14:17
package requests

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub-lesson/pkg/database"
	"gohub-lesson/pkg/response"
)

// init 初始化
func init() {
	// 注册自定义规则
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value any) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数，表名称，如 users
		tableName := rng[0]
		// 第二个参数，字段名称，如 email 或者 phone
		dbFiled := rng[1]

		// 第三个参数，排除 ID
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		// 用户请求过来的数据
		requestValue := value.(string)

		// 拼接 SQL
		query := database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue)

		// 如果传参第三个参数，加上 SQL Where 过滤
		if len(exceptID) > 0 {
			query.Where("id != ?", exceptID)
		}

		// 查询数据库
		var count int64
		query.Count(&count)

		// 验证不通过，数据库能找到对应的数据
		if count != 0 {
			// 如果有自定义错误消息的话
			if message != "" {
				return errors.New(message)
			}
			// 默认的错误消息
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		// 验证通过
		return nil
	})
}

type BaseRequest struct {
}

// MapErrs 返回数据
type MapErrs struct {
	url.Values
}

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

// ShouldBindJSON 解析数据
// request 引用(指针)
func ShouldBindJSON(request any, ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(request); err != nil {
		response.BadRequest(
			ctx,
			err,
			"请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
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
