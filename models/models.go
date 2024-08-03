// models.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Category  string `json:"category"`
	DueDate   time.Time
	Reminder  time.Time
}

type TodoMongo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Category  string             `bson:"category"`
	Completed bool               `bson:"completed"`
	DueDate   time.Time          `bson:"due_date"`
	Reminder  time.Time          `bson:"reminder"`
}
