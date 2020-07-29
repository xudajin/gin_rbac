package middleware

import (
	"go_web/util"

	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			util.Response(c, 403, 403, "请填写token", "")
			c.Abort()
			return
		}
		tokenClaims, err := util.ParseToken(token)
		if err != nil {
			util.Response(c, 403, 403, "token验证错误或超时", "")
			c.Abort()
			return
		}
		if tokenClaims != nil {
			// 以k-v 形式存储信息
			c.Set("username", tokenClaims.Username)
			c.Next()
		}
	}
}
