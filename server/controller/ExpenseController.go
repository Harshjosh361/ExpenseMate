package controller

import (
	"encoding/json"
	"net/http"
	"time"

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

func UpdateExpenseController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	// converting id
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var newExpense models.Expense
	err = json.NewDecoder(r.Body).Decode(&newExpense)
	if err != nil {
		http.Error(w, "invalid format", http.StatusBadRequest)
		return
	}

	// setting the id
	newExpense.Id = id

	res, err := models.UpdateExpense(newExpense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}

func FilterLastThreeMonthsController(w http.ResponseWriter, r *http.Request) {
	expenses, err := models.GetLastThreeMonths()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}
func FilterLastMonthControllerController(w http.ResponseWriter, r *http.Request) {
	expenses, err := models.GetLastMonth()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}
func FilterLastWeekController(w http.ResponseWriter, r *http.Request) {
	expenses, err := models.GetLastWeek()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}
func FilterByDateController(w http.ResponseWriter, r *http.Request) {
	startDateStr := r.URL.Query().Get("startdate")
	endDateStr := r.URL.Query().Get("enddate")

	// Parse Dates
	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	// Ensure the end date is after the start date
	if endDate.Before(startDate) {
		http.Error(w, "End date cannot be before start date", http.StatusBadRequest)
		return
	}
	expense, err := models.GetExpenseByDate(startDate, endDate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(expense); err != nil {
		http.Error(w, "failed to send data", http.StatusInternalServerError)
		return
	}
}
