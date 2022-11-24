// Package requests
// descr
// author fm
// date 2022/11/18 15:41
package requests

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"gohub-lesson/pkg/database"
)

// RuleNotExists 字段在数据库是否存在
func RuleNotExists(field string, rule string, message string, value any) error {
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
}

// RuleMaxCn 中文长度设定不大于
func RuleMaxCn(field string, rule string, message string, value any) error {
	// max_cn:8 中文长度设定不超过 8
	valLength := utf8.RuneCountInString(value.(string))

	l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))

	if valLength > l {
		// 如果有自定义错误消息的话，使用自定义消息
		if message != "" {
			return errors.New(message)
		}

		return fmt.Errorf("长度不能超过 %d 个字", l)
	}

	return nil

}

// RuleMinCn 中文长度设定不小于
func RuleMinCn(field string, rule string, message string, value any) error {
	// min_cn:2 中文长度设定不小于 2
	valLength := utf8.RuneCountInString(value.(string))

	l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))

	if valLength < l {
		// 如果有自定义错误消息的话，使用自定义消息
		if message != "" {
			return errors.New(message)
		}

		return fmt.Errorf("长度需大于 %d 个字", l)
	}

	return nil
}
