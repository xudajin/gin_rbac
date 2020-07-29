package router

import (
	"go_web/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userRouter := r.Group("/api/user")
	{
		userRouter.GET("/info/:user_id", controller.QueryUserById)
		userRouter.POST("/infos", controller.AddUser)
		userRouter.PUT("/info/:user_id", controller.UpdateById)
		userRouter.DELETE("/info/:user_id", controller.DeleteById)
	}
}
