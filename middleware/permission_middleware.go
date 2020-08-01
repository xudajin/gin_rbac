package middleware

import (
	"go_web/model"
	"go_web/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 权限验证中间件
func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestMethod := c.Request.Method // 获取当前请求方法
		requestURL := c.FullPath()        // 获取当前请求路径
		// 获取当前访问用户的权限
		name, ok := c.Get("username")
		if !ok {
			util.Response(c, http.StatusForbidden, 403, "访问失败", "")
		}
		var allowRequest bool = false // 定义标志变量
		// 判断是否是admin用户
		if name == "admin" {
			allowRequest = true
		} else {
			role, err := model.QueryPermissionByUserName(name.(string)) //类型断言
			if err != nil {
				util.Response(c, http.StatusForbidden, 403, "没有访问权限", "")
				c.Abort()
				return
			}
			for _, permission := range role.Permissions {
				if permission.Path == requestURL {
					if permission.Method == requestMethod {
						allowRequest = true
						break //有访问权限时，跳出循环
					}
				}
			}
		}
		// 判断是否允许访问
		if allowRequest {
			c.Next()
		} else {
			util.Response(c, http.StatusForbidden, 403, "没有访问权限", "")
			c.Abort()
			return
		}
	}
}
