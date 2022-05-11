package controllers

import (
	"TheWarEconomy/api/models"
	"encoding/json"
	"net/http"
)

type UserResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Error    string `json:"error"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&UserResponse{
			Error: "Unable to sign up, please try again",
		})
		return
	}

	u, err := user.Create()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&UserResponse{
			Error: "Unable to sign up, please try again",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&UserResponse{
		Id:       u.Id,
		Username: u.Username,
		Email:    u.Email,
	})
}
