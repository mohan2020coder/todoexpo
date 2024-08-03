// main.go
package main

import (
	"context"
	"log"
	"todoexpo/config"
	"todoexpo/controllers"
	"todoexpo/models"
	"todoexpo/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbType           string
	mongoClient      *mongo.Client
	mongoCollection  *mongo.Collection
	sqliteDB         *gorm.DB
	sqliteController *controllers.SQLiteController
	mongoController  *controllers.MongoController
)

func initDB() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	dbType = config.Database.Type

	if dbType == "mongodb" {
		mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Database.MongoDB.URI))
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}
		err = mongoClient.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal("Failed to ping MongoDB:", err)
		}
		mongoCollection = mongoClient.Database(config.Database.MongoDB.Database).Collection("todos")
		mongoController = &controllers.MongoController{Collection: mongoCollection}
	} else if dbType == "sqlite" {
		sqliteDB, err = gorm.Open(sqlite.Open(config.Database.SQLite.Path), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to SQLite database:", err)
		}
		// Define your models and auto-migrate here
		sqliteDB.AutoMigrate(&models.Todo{})
		sqliteController = &controllers.SQLiteController{DB: sqliteDB}
	} else {
		log.Fatal("Unknown database type:", dbType)
	}
}

func main() {
	initDB()

	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // List allowed origins here
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Set up routes based on the database type
	routes.SetupRoutes(r, dbType, mongoController, sqliteController)

	r.Run(":8080")
}
