package router

import (
	"go_web/controller"
	"go_web/middleware"

	"github.com/gin-gonic/gin"
)

func PermissionRouter(r *gin.Engine) {
	// 设置中间件
	permissionRouter := r.Group("/api/permission", middleware.Jwt(), middleware.Permission())
	{
		permissionRouter.GET("/infos", controller.PermissionList)
		permissionRouter.POST("/infos", controller.AddPermission)
		permissionRouter.PUT("/info/:permission_id", controller.UpdatePermission)
		permissionRouter.DELETE("/info/:permission_id", controller.DeletePermission)
	}
}
