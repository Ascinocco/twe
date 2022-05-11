package models

import (
	"TheWarEconomy/api/database"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                   string `json:"id"`
	Email                string `json:"email"`
	Username             string `json:"username"`
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

	if len(user.Username) < 3 {
		return "Username must be greater than 3 characters.", false
	}

	return "", true
}

func (user *User) Create() (*User, error) {
	if errMsg, ok := user.Validate(); !ok {
		return user, errors.New(errMsg)
	}

	hashedPw, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Id = user.Email
	user.Password = string(hashedPw)
	user.PasswordConfirmation = ""

	uc := database.GetUsersCol()

	// @TODO: How to insure email, username and PMC name are unique? use look up document, make the email the id...
	// @TODO: don't leak email addresses?
	_, err := uc.Insert(user.Id, user, nil)

	return user, err
}
