package util

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, status int, code int64, msg string, data interface{}) {
	c.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ListResponse(c *gin.Context, status int, code int64, msg string, count int, data interface{}) {
	c.JSON(status, gin.H{
		"code":  code,
		"msg":   msg,
		"data":  data,
		"count": count,
	})
}
