package util

import (
	"github.com/gin-gonic/gin"
)

// 常规返回
func Response(c *gin.Context, status int, code int64, msg string, data interface{}) {
	c.JSON(status, gin.H{
		"code":   code,
		"msg":    msg,
		"result": data,
	})
}

// 列表数据返回
func ListResponse(c *gin.Context, status int, code int64, msg string, count int, data interface{}) {
	c.JSON(status, gin.H{
		"code":   code,
		"msg":    msg,
		"result": data,
		"count":  count,
	})
}
