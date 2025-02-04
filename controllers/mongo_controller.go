// mongo_controller.go
package controllers

import (
	"context"
	"net/http"
	"todoexpo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoController struct {
	Collection *mongo.Collection
}

func (ctrl *MongoController) GetTodos(c *gin.Context) {
	// var todos []bson.M
	var todos []models.TodoMongo
	cursor, err := ctrl.Collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = cursor.All(context.TODO(), &todos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (ctrl *MongoController) CreateTodo(c *gin.Context) {
	// var todo bson.M
	var todo models.TodoMongo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := ctrl.Collection.InsertOne(context.TODO(), todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (ctrl *MongoController) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	//var todo bson.M
	var todo models.TodoMongo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := ctrl.Collection.UpdateByID(context.TODO(), id, bson.M{"$set": todo})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (ctrl *MongoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	_, err := ctrl.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

func (ctrl *MongoController) GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	//var todo bson.M
	var todo models.TodoMongo
	err := ctrl.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}
func (controller *MongoController) UpdateTodoByID(c *gin.Context) {
	id := c.Param("id")
	var todo models.TodoMongo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": todo}
	_, err = controller.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}
