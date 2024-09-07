package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Harshjosh361/ExpenseMate/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategoryController(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}
	// add id to category
	newCategory.Id = primitive.NewObjectID()

	// invoking create category to interact with database
	err := models.CreateCategory(&newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "category created successfully",
		"_id":     newCategory.Id.Hex(),
		"name":    newCategory.Name,
	})
}
func GetCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetAllCategory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// sending categories
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, "Failed to encode categories", http.StatusInternalServerError)
	}
}

func GetSingleCategoryController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category, err := models.GetSingleCategory(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// sending categories
	if err := json.NewEncoder(w).Encode(category); err != nil {
		http.Error(w, "Failed to encode categories", http.StatusInternalServerError)
	}

}

func DeleteCategoryController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := models.DeleteCategory(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "deleted successfully",
		"_id":     params["id"],
	})
}
