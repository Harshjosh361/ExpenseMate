package models

import (
	"context"
	"errors"

	"github.com/Harshjosh361/ExpenseMate/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Category struct {
	Id   primitive.ObjectID `bson:"_id" json:"_id"`
	Name string             `bson:"name" json:"name"`
}

func CreateCategory(category *Category) error {
	// Check for duplicate category by name
	var existingCategory Category
	err := db.CollectionCategory.FindOne(context.Background(), bson.M{"name": category.Name}).Decode(&existingCategory)

	if err == nil {
		// If no error and category is found, it means the category exists
		return errors.New("category already exists")
	} else if err != mongo.ErrNoDocuments {
		// If it's any other error, return it
		// ErrNoDocuments is return by decode method if np documnets found based on filter
		return errors.New("failed to check for existing category")
	}

	// If no duplicate, proceed to insert the new category
	_, err = db.CollectionCategory.InsertOne(context.Background(), category)
	if err != nil {
		return errors.New("failed to create new category")
	}

	return nil
}

func GetAllCategory() ([]Category, error) {
	Cursor, err := db.CollectionCategory.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("failed to fetch categories")
	}
	defer Cursor.Close(context.Background())
	// slice to store fetched categories
	var categories []Category

	for Cursor.Next(context.Background()) {
		var category Category
		if err := Cursor.Decode(&category); err != nil {
			return nil, errors.New("failed to decode category")
		}
		categories = append(categories, category)
	}
	return categories, nil
}
func GetSingleCategory(id string) (*Category, error) {
	// convert the id from string to ObjectID
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid category id")
	}

	var category Category
	err = db.CollectionCategory.FindOne(context.Background(), bson.M{"_id": ObjectID}).Decode(&category)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("category not found")
	} else if err != nil {
		return nil, errors.New("failed to fetch category")
	}
	return &category, nil
}

func DeleteCategory(id string) error {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid category id")
	}

	result, err := db.CollectionCategory.DeleteOne(context.Background(), bson.M{"_id": ObjectID})
	if err != nil {
		return errors.New("failed to delete category")
	}
	if result.DeletedCount == 0 {
		return errors.New("Category not found")
	}
	return nil
}
