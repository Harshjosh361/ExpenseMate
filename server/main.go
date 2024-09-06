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

	// Initializing AuthRoutes using SubRouter
	routes.AuthRoute(authrouter)

	log.Fatal(http.ListenAndServe(":8000", r))
}

// Testing home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Server is live")
	if err != nil {
		log.Fatal(err)
	}
}
