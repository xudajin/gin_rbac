package main

import (
	// 初始化配置文件

	"go_web/config"
	"go_web/controller"
	_ "go_web/model"
	"go_web/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Set.Mode) // 设置gin的模式运行，默认Debug
	// 根路由
	route := gin.Default()
	route.POST("admin/login", controller.LoginController)
	//注册admin管理路由
	router.AdminRouter(route)

	route.Run(":8080")

}
