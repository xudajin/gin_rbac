package router

import (
	"go_web/controller"

	"github.com/gin-gonic/gin"
)

func RoleRouter(r *gin.Engine) {
	roleRouter := r.Group("/api/role")
	{
		roleRouter.GET("/infos", controller.QueryRoles)
		roleRouter.POST("/infos", controller.AddRole)
		roleRouter.PUT("/info/:role_id", controller.UpdateRole)
		roleRouter.DELETE("/info/:role_id", controller.DeleteRole)
		roleRouter.POST("/permissions/:role_id", controller.RoleAddPermission)
	}
}
