package routes

import (
	"github.com/Harshjosh361/ExpenseMate/controller"
	"github.com/gorilla/mux"
)

func ExpenseRoute(r *mux.Router) {
	r.HandleFunc("/create-expense", controller.CreateExpenseController).Methods("POST")
	r.HandleFunc("/get-expense", controller.GetAllExpenseController).Methods("GET")
	r.HandleFunc("/get-expense/{id}", controller.GetExpenseController).Methods("GET")
	r.HandleFunc("/delete-expense/{id}", controller.DeleteExpenseController).Methods("DELETE")
}
