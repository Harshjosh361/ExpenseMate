package models

import (
	"context"
	"errors"
	"time"

	"github.com/Harshjosh361/ExpenseMate/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetAllExpense() ([]Expense, error) {
	Cursor, err := db.CollectionExpense.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("failed to find expenses")
	}
	defer Cursor.Close(context.Background())

	var expenses []Expense
	for Cursor.Next(context.Background()) {
		var expense Expense
		err = Cursor.Decode(&expense)
		if err != nil {
			return nil, errors.New("error in decoding")
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil

}

func GetExpense(id string) (Expense, error) {
	// convert the id from string to ObjectID
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Expense{}, errors.New("invalid expense id")
	}

	var expense Expense
	err = db.CollectionExpense.FindOne(context.Background(), bson.M{"_id": ObjectID}).Decode(&expense)
	if err == mongo.ErrNoDocuments {
		return Expense{}, errors.New("no expense found")
	} else if err != nil {
		return Expense{}, err
	}
	return expense, nil
}

func DeleteExpense(id string) error {
	// convert the id from string to ObjectID
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid expense id")
	}
	res, _ := db.CollectionExpense.DeleteOne(context.Background(), bson.M{"_id": ObjectID})
	if res.DeletedCount == 0 {
		return errors.New("failed to delete expense")
	}
	return nil
}

func UpdateExpense(newExpense Expense) (Expense, error) {

	var result Expense
	update := bson.M{"$set": bson.M{
		"title":    newExpense.Title,
		"amount":   newExpense.Amount,
		"category": newExpense.Category,
		"date":     time.Now(),
	}}
	// db.CollectionExpense.UpdateOne(context.Background(), bson.M{}, update)
	// FindOneAndUpdate gives back the document wheras updateone just give the count
	err := db.CollectionExpense.FindOneAndUpdate(context.Background(), bson.M{}, update).Decode(&result)
	if err != nil {
		return Expense{}, errors.New("failed to update")
	}
	return result, nil
}

func GetLastThreeMonths() ([]Expense, error) {
	threeMonthsAgo := time.Now().AddDate(0, -3, 0)
	filter := bson.M{
		"date": bson.M{"$gt": threeMonthsAgo},
	}

	Cursor, err := db.CollectionExpense.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.New("failed to get data")
	}
	defer Cursor.Close(context.Background())

	var expenses []Expense
	for Cursor.Next(context.Background()) {
		var expense Expense
		Cursor.Decode(&expense)
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func GetLastMonth() ([]Expense, error) {
	lastMonth := time.Now().AddDate(0, -1, 0)
	filter := bson.M{
		"date": bson.M{"$gt": lastMonth},
	}

	Cursor, err := db.CollectionExpense.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.New("failed to get data")
	}
	defer Cursor.Close(context.Background())

	var expenses []Expense
	for Cursor.Next(context.Background()) {
		var expense Expense
		Cursor.Decode(&expense)
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
func GetLastWeek() ([]Expense, error) {
	lastWeek := time.Now().AddDate(0, 0, -7)
	filter := bson.M{
		"date": bson.M{"$gt": lastWeek},
	}

	Cursor, err := db.CollectionExpense.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.New("failed to get data")
	}
	defer Cursor.Close(context.Background())

	var expenses []Expense
	for Cursor.Next(context.Background()) {
		var expense Expense
		Cursor.Decode(&expense)
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func GetExpenseByDate(startdate, endDate time.Time) ([]Expense, error) {
	filter := bson.M{
		"date": bson.M{"$gt": startdate, "$lt": endDate},
	}

	Cursor, err := db.CollectionExpense.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.New("failed to get data")
	}
	defer Cursor.Close(context.Background())

	var expenses []Expense
	for Cursor.Next(context.Background()) {
		var expense Expense
		Cursor.Decode(&expense)
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
