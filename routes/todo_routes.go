package routes

import (
	"awesomeProject/controllers"
	"github.com/gin-gonic/gin"
)

// 注册 Todo 路由
func RegisterTodoRoutes(r *gin.Engine) {
	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
}
