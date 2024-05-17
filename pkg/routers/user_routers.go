package routers

import (
	"encoding/json"
	util "gorl/pkg"
	"gorl/pkg/db"
	"net/http"
)

type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func CreateUserRouter(w http.ResponseWriter, r *http.Request) {
	var request createUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed parsing Request", http.StatusBadRequest)
		return
	}

	err = db.CreateUser(request.Username, request.Password, request.IsAdmin)
	if err != nil {
		http.Error(w, "User already exist", http.StatusBadRequest)
		return
	}

	token, err := util.SignedJwt(request.Username, request.IsAdmin)
	if err != nil {
		http.Error(w, "Failed to create JWT token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("TOKEN", token)
	w.WriteHeader(http.StatusCreated)
}
