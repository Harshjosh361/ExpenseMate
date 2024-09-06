package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Title    string             `bson:"title" json:"title"`
	Amount   float64            `bson:"amount" json:"amount"`
	Category primitive.ObjectID `bson:"category" json:"category"`
	Date     time.Time          `bson:"date" json:"date"`
}
