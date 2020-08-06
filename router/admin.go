package router

import (
	"go_web/controller"
	"go_web/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	adminRouter := r.Group("/admin", middleware.Jwt(), middleware.Permission())
	// 用户路由
	{
		adminRouter.GET("/users", controller.QueryUserList)
		adminRouter.POST("/users", controller.AddUser)
		adminRouter.PUT("/user/:user_id", controller.UpdateUser)
		adminRouter.PUT("/password/:user_id", controller.ChangePassword)
		adminRouter.DELETE("/user/:user_id", controller.DeleteUser)
	}
	// 角色路由
	{
		adminRouter.GET("/roles", controller.QueryRoles)
		adminRouter.POST("/roles", controller.AddRole)
		adminRouter.PUT("/roles/:role_id", controller.UpdateRole)
		adminRouter.DELETE("/roles/:role_id", controller.DeleteRole)
		adminRouter.GET("role/permissions/:role_id", controller.QueryPermissionByRoleID)
		adminRouter.POST("role/permissions/:role_id", controller.RoleAddPermission)
	}
	// 权限路由
	{
		adminRouter.GET("/permissions", controller.PermissionList)
		adminRouter.POST("/permissions", controller.AddPermission)
		adminRouter.PUT("/permission/:permission_id", controller.UpdatePermission)
		adminRouter.DELETE("/permission/:permission_id", controller.DeletePermission)
	}
}
