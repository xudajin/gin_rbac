package router

import (
	"go_web/controller"

	"github.com/gin-gonic/gin"
)

func PermissionRouter(r *gin.Engine) {
	permissionRouter := r.Group("/api/permission")
	{
		permissionRouter.GET("/infos", controller.PermissionList)
		permissionRouter.POST("/infos", controller.AddPermission)
		permissionRouter.PUT("/info/:permission_id", controller.UpdatePermission)
	}
}
