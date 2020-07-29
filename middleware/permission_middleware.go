package middleware

import (
	"go_web/model"
	"go_web/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestMethod := c.Request.Method // 获取当前请求方法
		requestURL := c.FullPath()        // 获取当前请求路径
		// 获取当前访问用户的权限
		name, ok := c.Get("username")
		if !ok {
			util.Response(c, http.StatusForbidden, 403, "访问失败", "")
		}

		role, err := model.QueryPermissionByUserName(name.(string)) //类型断言
		if err != nil {
			util.Response(c, http.StatusForbidden, 403, "没有访问权限", "")
			c.Abort()
			return
		}
		// 判断权限是否通过
		var allowRequest bool = false
		for _, permission := range role.Permissions {
			if permission.Path == requestURL {
				if permission.Method == requestMethod {
					allowRequest = true
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
