package http

import (
	"clean-polytech/internal/app/user"
	"clean-polytech/internal/infra/db/postgres"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	saveUserUC *user.SaveUser
	getUsersUC *user.GetUsersUse
}

func (h *UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewUserHandler(saveUserUC *postgres.UserRepository, getUsersUC *postgres.PhoneRepository) *UserHandler {
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
