// routes.go
package routes

import (
	"todoexpo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, dbType string, mongoController *controllers.MongoController, sqliteController *controllers.SQLiteController) {
	if dbType == "mongodb" {
		r.GET("/todos", mongoController.GetTodos)
		r.POST("/todos", mongoController.CreateTodo)
		r.PUT("/todos/:id", mongoController.UpdateTodo)
		r.DELETE("/todos/:id", mongoController.DeleteTodo)
		r.GET("/todos/:id", mongoController.GetTodoByID)
	} else if dbType == "sqlite" {
		r.GET("/todos", sqliteController.GetTodos)
		r.POST("/todos", sqliteController.CreateTodo)
		r.PUT("/todos/:id", sqliteController.UpdateTodo)
		r.DELETE("/todos/:id", sqliteController.DeleteTodo)
		r.GET("/todos/:id", sqliteController.GetTodoByID)
	} else {
		panic("Unknown database type: " + dbType)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})
}
