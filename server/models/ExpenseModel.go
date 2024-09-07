package models

import (
	"context"
	"errors"
	"time"

	"github.com/Harshjosh361/ExpenseMate/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Title    string             `bson:"title" json:"title"`
	Amount   float64            `bson:"amount" json:"amount"`
	Category primitive.ObjectID `bson:"category" json:"category"`
	Date     time.Time          `bson:"date" json:"date"`
}

func CreateExpense(expense *Expense) (Expense, error) {
	// Check for existing expense title
	var existingExpense Expense
	err := db.CollectionExpense.FindOne(context.Background(), bson.M{"title": expense.Title}).Decode(&existingExpense)
	if err == nil {
		return Expense{}, errors.New("expense already exist")
	}

	//create expense
	_, err = db.CollectionExpense.InsertOne(context.Background(), expense)
	if err != nil {
		return Expense{}, errors.New("failed to create expense")
	}
	return *expense, nil
}
