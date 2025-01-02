package controllers

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var todos = []models.Todo{
	{ID: 1, Title: "Learn Go", Completed: false},
	{ID: 2, Title: "Build a web app", Completed: false},
}

// 获取所有 Todo 列表
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// 创建新的 Todo
func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 生成 ID
	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

// 更新 Todo
func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查找并更新 Todo
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updatedTodo
			todos[i].ID = id // 保持 ID 不变
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// 删除 Todo
func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, todo := range todos {
		if todo.ID == id {
			// 删除 Todo
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
