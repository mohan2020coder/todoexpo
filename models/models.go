// models.go
package models

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title"`
	Completed bool       `json:"completed"`
	Category  string     `json:"category"`
	DueDate   CustomTime `json:"due_date"`
	Reminder  CustomTime `json:"reminder"`
	Priority  string     `json:"priority"`
}

type TodoMongo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Category  string             `bson:"category"`
	Completed bool               `bson:"completed"`
	DueDate   time.Time          `bson:"due_date"`
	Reminder  time.Time          `bson:"reminder"`
	Priority  string             `bson:"priority"` // Add priority field
}

const (
	customTimeFormat = "2006-01-02T15:04"
)

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	t, err := time.Parse(`"`+customTimeFormat+`"`, str)
	if err != nil {
		return fmt.Errorf("parsing time %s as %s: %w", str, customTimeFormat, err)
	}
	*ct = CustomTime(t)
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ct).Format(customTimeFormat))
}
