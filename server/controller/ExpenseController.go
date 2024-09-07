package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Harshjosh361/ExpenseMate/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateExpenseController(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "invalid format", http.StatusBadRequest)
		return
	}

	// generate new id
	expense.Id = primitive.NewObjectID()

	res, err := models.CreateExpense(&expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "expense created",
		"_id":     res.Id.Hex(),
	})

}
