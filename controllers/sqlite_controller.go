// sqlite_controller.go
package controllers

import (
	"net/http"
	"todoexpo/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SQLiteController struct {
	DB *gorm.DB
}

func (ctrl *SQLiteController) GetTodos(c *gin.Context) {
	var todos []models.Todo
	ctrl.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func (ctrl *SQLiteController) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func (ctrl *SQLiteController) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := ctrl.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func (ctrl *SQLiteController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

func (ctrl *SQLiteController) GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := ctrl.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}
