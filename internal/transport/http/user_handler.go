package http

import (
	"clean-polytech/internal/app/user"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	saveUserUC *user.SaveUser
	getUsersUC *user.GetUsersUse
}

// Implement the ServeHTTP method for UserHandler
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// For simplicity, we handle only two basic routes here
	switch r.URL.Path {
	case "/save_users":
		// Call the use case for saving a user (example, modify as needed)
		ctx := context.Background()
		err := h.saveUserUC.Execute(ctx)
		if err != nil {
			http.Error(w, "Failed to save user", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "User saved successfully")

	case "/get":
		// Call the use case for getting users (example, modify as needed)
		users, err := h.getUsersUC.Execute()
		if err != nil {
			http.Error(w, "Failed to get users", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Users: %v", users)

	default:
		http.NotFound(w, r)
	}
}
func NewUserHandler(saveUserUC *user.SaveUser, getUsersUC *user.GetUsersUse) *UserHandler {
	return &UserHandler{
		saveUserUC: saveUserUC,
		getUsersUC: getUsersUC,
	}
}

func (h *UserHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name       string `json:"name"`
		FamilyName string `json:"family_name"`
		Charge     string `json:"charge"`
		PhoneType  string `json:"phone_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.saveUserUC.Execute(r.Context(), req.Name, req.FamilyName, req.PhoneType, req.Charge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.getUsersUC.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
