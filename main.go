package main

import (
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 Gin 引擎
	r := gin.Default()

	// 注册路由
	routes.RegisterTodoRoutes(r)

	// 启动服务
	r.Run(":8080")
}
