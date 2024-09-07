package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Harshjosh361/ExpenseMate/helper"
	"github.com/Harshjosh361/ExpenseMate/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	// Set the userId
	user.Id = primitive.NewObjectID()

	// call the registermethod
	err := models.RegisterUser(&user)
	if err != nil {
		if err.Error() == "User already exists" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, "Error registering user", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered Successfully"})

}

func LoginController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}
	foundUser, err := models.LoginUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// generate JWT token
	token, err := helper.GenerateJWT(foundUser.Id.Hex())
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Sending Response with token and user data
	// no need to write header json will automatically give status 200 OK
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login Successfull",
		"token":   token,
		"user":    foundUser.Name,
	})
}
