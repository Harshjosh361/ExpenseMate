package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Harshjosh361/ExpenseMate/db"
	"github.com/Harshjosh361/ExpenseMate/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.ConnectDb()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	authrouter := r.PathPrefix("/api/v1/auth").Subrouter()
	categoryrouter := r.PathPrefix("/api/v1/category").Subrouter()
	expenserouter := r.PathPrefix("/api/v1/expense").Subrouter()

	// Initializing AuthRoutes using SubRouter
	routes.AuthRoute(authrouter)
	routes.CateogryRoute(categoryrouter)
	routes.ExpenseRoute(expenserouter)

	log.Fatal(http.ListenAndServe(":8000", r))
}

// Testing home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Server is live")
	if err != nil {
		log.Fatal(err)
	}
}
