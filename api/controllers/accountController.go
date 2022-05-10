package controllers

import (
	"TheWarEconomy/api/database"
	"TheWarEconomy/api/models"
	"TheWarEconomy/api/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/couchbase/gocb/v2"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PmcResponse struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

type SignUpResponse struct {
	User  UserResponse
	Pmc   PmcResponse
	Token string `json:"token"`
	Error string `json:"error"`
}

type Token struct {
	UserId string
	jwt.StandardClaims
}

func authenticate(email, password string) (string, bool) {
	user := &models.User{}
	twec := database.GetTweCluster()
	qr, err := twec.Query("SELECT * FROM `users` WHERE email = $1", &gocb.QueryOptions{PositionalParameters: []interface{}{email}})

	if err != nil {
		return "Could not sign in.", false
	}

	err = qr.One(user)

	if err != nil {
		return "Could not sign in", false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "Could not sign in.", false
	}

	tokenData := &Token{UserId: user.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenData)
	tokenString, err := token.SignedString([]byte(os.Getenv(utils.EnvTokenSecret)))

	if err != nil {
		return "Could not sign in", false
	}

	return tokenString, true
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	pmc := &models.Pmc{
		Name: user.PmcName,
	}

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&SignUpResponse{
			Error: "Json decoding failed.",
		})
		return
	}

	p, err := pmc.Create()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&SignUpResponse{
			Error: "PMC creation failed.",
		})
		return
	}

	u, err := user.Create(p.Id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&SignUpResponse{
			Error: "User creation failed.",
		})
		return
	}

	_, err = pmc.LinkUser(u.Id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&SignUpResponse{
			Error: "PMC to User association failed.",
		})
		return
	}

	// res = tokenString or err message... probably bad...
	// @TODO: separate results?
	tokenRes, ok := authenticate(user.Email, user.Password)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&SignUpResponse{
			Error: tokenRes,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SignUpResponse{
		User: UserResponse{
			Id:       u.Id,
			Username: u.Username,
			Email:    u.Email,
		},
		Pmc: PmcResponse{
			Id:     p.Id,
			UserId: u.Id,
			Name:   p.Name,
		},
		Token: tokenRes,
		Error: "",
	})
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}
