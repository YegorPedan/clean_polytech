package http

import (
	"clean-polytech/internal/app/user"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	saveUserUC *user.SaveUser
	getUsersUC *user.GetUsersUse
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
		Phone      string `json:"phone"`
		PhoneType  string `json:"phone_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.saveUserUC.Execute(r.Context(), req.Name, req.FamilyName, req.Phone)
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
