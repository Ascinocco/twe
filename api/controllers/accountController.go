package controllers

import (
	"TheWarEconomy/api/database"
	"TheWarEconomy/api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
)

// @TODO: Handle responses, remove os.exit.
// @TODO: Db indexes? Auto incrementing ids?
func SignUp(w http.ResponseWriter, r *http.Request) {
	uc := database.GetUsersCol()
	id := uuid.NewString()
	nu := models.User{
		Id:       id,
		Name:     "Test",
		Email:    "test@mail.com",
		Password: "test",
	}

	_, err := uc.Upsert(id, nu, nil)

	if err != nil {
		fmt.Println("Couldn't create user", err)
		// swap for error response
		os.Exit(1)
	}

	getResult, err := uc.Get(id, nil)

	if err != nil {
		fmt.Println("Couldn't get user", err)
		os.Exit(1)
	}

	var foundUser models.User
	err = getResult.Content(&foundUser)

	if err != nil {
		fmt.Println("Couldn't parse user result", err)
		os.Exit(1)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundUser)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}
