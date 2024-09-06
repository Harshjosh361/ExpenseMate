package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Testing home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Server is live")
	if err != nil {
		log.Fatal(err)
	}
}
