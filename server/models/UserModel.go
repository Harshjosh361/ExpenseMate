package models

import (
	"context"
	"errors"

	"github.com/Harshjosh361/ExpenseMate/db"
	"github.com/Harshjosh361/ExpenseMate/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

func RegisterUser(user *User) error {
	var existingUser User
	err := db.CollectionUser.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("User already exists")
	}
	// hashing password before inserting
	user.Password, _ = helper.HashPassword(user.Password)

	_, err = db.CollectionUser.InsertOne(context.Background(), user)
	return err
}

func LoginUser(user *User) (*User, error) {
	var foundUser User
	err := db.CollectionUser.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check password
	if err := helper.Checkpassword(foundUser.Password, user.Password); err != nil {
		return nil, errors.New("invalid  password")
	}
	return &foundUser, nil
}
