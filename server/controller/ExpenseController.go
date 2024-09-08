package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Harshjosh361/ExpenseMate/models"
	"github.com/gorilla/mux"
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

func GetAllExpenseController(w http.ResponseWriter, r *http.Request) {
	expenses, err := models.GetAllExpense()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(expenses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetExpenseController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	expense, err := models.GetExpense(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(expense); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
	}
}

func DeleteExpenseController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := models.DeleteExpense(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "expense deleted successfully",
	})
}
