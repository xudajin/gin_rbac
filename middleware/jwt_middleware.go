package middleware

import (
	"go_web/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// jwt token验证
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			util.Response(c, http.StatusUnauthorized, 403, "请填写token", "")
			c.Abort()
			return
		}
		tokenClaims, err := util.ParseToken(token)
		if err != nil {
			util.Response(c, http.StatusUnauthorized, 403, "token验证错误或超时", "")
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
