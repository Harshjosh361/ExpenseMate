package routes

import (
	"github.com/Harshjosh361/ExpenseMate/controller"
	"github.com/gorilla/mux"
)

func AuthRoute(r *mux.Router) {
	// r.HandleFunc("/login", controller.LoginHandler()).Methods("POST")
	r.HandleFunc("/register", controller.RegisterController).Methods("POST")
	r.HandleFunc("/login", controller.LoginController).Methods("POST")
}
