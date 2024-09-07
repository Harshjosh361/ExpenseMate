package routes

import (
	"github.com/Harshjosh361/ExpenseMate/controller"
	"github.com/gorilla/mux"
)

func ExpenseRoute(r *mux.Router) {
	r.HandleFunc("/create-expense", controller.CreateExpenseController).Methods("POST")
	// r.HandleFunc("/get-category", controller.GetCategory).Methods("GET")
	// r.HandleFunc("/get-category/{id}", controller.GetSingleCategoryController).Methods("GET")
	// r.HandleFunc("/delete-category/{id}", controller.DeleteCategoryController).Methods("DELETE")
}
