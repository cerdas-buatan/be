package module

import (
	"encoding/json"
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/cerdas-buatan/be/helper"
	"github.com/cerdas-buatan/be/model"
)

// HomeGaysdisal handles the home endpoint
func HomeGaysdisal(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Gaysdisal AI %s", "8080")
	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error: JSON marshaling failed", http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

// RenameMenuHandler handles renaming a menu
func RenameMenuHandler(s *model.MenuService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			ID      string `json:"id"`
			NewName string `json:"new_name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := primitive.ObjectIDFromHex(request.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := s.RenameMenu(r.Context(), id, request.NewName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(model.Response{Status: true, Message: "Menu renamed successfully"})
	}
}

// ArchiveMenuHandler handles moving a menu to the archive
func ArchiveMenuHandler(s *model.MenuService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			ID string `json:"id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := primitive.ObjectIDFromHex(request.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := s.ArchiveMenu(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(model.Response{Status: true, Message: "Menu moved to archive successfully"})
	}
}

// AddMenuHandler handles adding a new menu
func AddMenuHandler(s *model.MenuService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var menu model.Menu
		if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		menu, err := s.AddMenu(r.Context(), menu)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(model.Response{Status: true, Message: "New menu added successfully", Data: menu})
	}
}

// NotFound handles 404 errors
func NotFound(w http.ResponseWriter, r *http.Request) {
	var resp model.Response
	resp.Message = "Not Found"
	helper.WriteJSON(w, http.StatusNotFound, resp)
}
