package models

import (
	"TheWarEconomy/api/database"
	"errors"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                   string `json:"id"`
	Email                string `json:"email"`
	Username             string `json:"username"`
	PmcName              string `json:"pmcName"`
	PmcId                string `json:"pmcId"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

func (user *User) Validate() (string, bool) {
	if !strings.Contains(user.Email, "@") {
		return "Invalid email.", false
	}

	// @TODO: Improve password validation
	if len(user.Password) < 6 {
		return "Password must include numbers, letters and symbols.", false
	}

	if strings.Compare(user.Password, user.PasswordConfirmation) != 0 {
		return "Passwords do not match.", false
	}

	if len(user.PmcName) <= 5 {
		return "PMC name must be at least 5 characters", false
	}

	if len(user.Username) < 3 {
		return "Username must be greater than 3 characters.", false
	}

	return "", true
}

func (user *User) Create(pmcId string) (*User, error) {
	if errMsg, ok := user.Validate(); !ok {
		return user, errors.New(errMsg)
	}

	hashedPw, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Id = uuid.NewString()
	user.Password = string(hashedPw)
	user.PasswordConfirmation = ""
	user.PmcId = pmcId

	uc := database.GetUsersCol()

	_, err := uc.Insert(user.Id, user, nil)

	return user, err
}