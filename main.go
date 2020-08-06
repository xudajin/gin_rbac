package main

import (
	"fmt"
	_ "go_web/config"
	"go_web/controller"
	"go_web/model"
	"go_web/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode) // 设置gin的模式运行，默认Debug
	// 根路由
	route := gin.Default()
	route.POST("admin/login", controller.LoginController)
	//后台管理路由
	router.AdminRouter(route)
	fmt.Println(model.DB)

	route.Run(":8080")

}
