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
	r.HandleFunc("/update-expense/{id}", controller.UpdateExpenseController).Methods("PUT")
	r.HandleFunc("/filter-expense/lastweek", controller.FilterLastWeekController).Methods("GET")
	r.HandleFunc("/filter-expense/lastmonth", controller.FilterLastMonthControllerController).Methods("GET")
	r.HandleFunc("/filters-expense/lastthreemonths", controller.FilterLastThreeMonthsController).Methods("GET")
	r.HandleFunc("/filter-expense", controller.FilterByDateController).Methods("GET")
}
