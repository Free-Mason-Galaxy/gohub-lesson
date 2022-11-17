// Package response
// descr 带自定义状态码响应处理工具
//
//	自定义状态码可以用于快速定位，如：一个URI多个地方返回同一错误信息，使用不同的自定义状态来区分
//	如：因特殊原因不管成功错误与否 httpStatus 码都是返回 200 时
//
// author fm
// date 2022/11/17 11:17
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CJSON 响应 200 和 JSON 数据
func CJSON(c *gin.Context, status int, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": "",
		"data":    defaultData(data),
	})
}

// defaultData 内用的辅助函数，用以支持默认参数默认值
func defaultData(data gin.H) gin.H {
	if data == nil {
		return make(gin.H)
	}
	return data
}
