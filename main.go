package main

import (
	"fmt"
	"go_web/controller"
	"go_web/model"
	"go_web/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 根路由
	route := gin.Default()
	route.POST("/login", controller.LoginController)

	// app路由
	router.UserRouter(route)
	router.PermissionRouter(route)
	router.RoleRouter(route)
	// 用户模块路由
	fmt.Println(model.DB)

	route.Run(":8080")

}
