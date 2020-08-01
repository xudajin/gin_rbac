package router

import (
	"go_web/controller"
	"go_web/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userRouter := r.Group("/api/user", middleware.Jwt(), middleware.Permission())
	{
		userRouter.GET("/infos", controller.QueryUserList)
		userRouter.GET("/info/:user_id", controller.QueryUser)
		userRouter.POST("/infos", controller.AddUser)
		userRouter.PUT("/info/:user_id", controller.UpdateUser)
		userRouter.PUT("/password/:user_id", controller.ChangePassword)
		userRouter.DELETE("/info/:user_id", controller.DeleteUser)
	}
}
