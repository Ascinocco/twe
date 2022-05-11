package controllers

import (
	"TheWarEconomy/api/database"
	"TheWarEconomy/api/models"
	"TheWarEconomy/api/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

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

func authenticate(email, password string) (string, error) {
	user := models.User{}
	res, err := database.GetUsersCol().Get(email, nil)

	// start - @TODO: Leave this for an example on querying...
	// twes := database.GetTweScope()

	// qr, err := twes.Query("SELECT U.* FROM `users` U WHERE email = $email", &gocb.QueryOptions{
	// 	NamedParameters: map[string]interface{}{
	// 		"email": email,
	// 	},
	// })

	// user := models.User{}
	// 	err = qr.One(&user)
	// end -

	res.Content(&user)

	if err != nil {
		fmt.Println("a", err)
		return "", errors.New("Could not sign in.")
	}

	if err != nil {
		fmt.Println("b", err)
		return "", errors.New("Could not sign in.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		fmt.Println("c", err)
		return "", errors.New("Could not sign in.")
	}

	tokenData := &Token{UserId: user.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenData)
	tokenString, err := token.SignedString([]byte(os.Getenv(utils.EnvTokenSecret)))

	if err != nil {
		fmt.Println("d", err)
		return "", errors.New("Could not sign in.")
	}

	return tokenString, nil
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&SignUpResponse{
			Error: "Json decoding failed.",
		})
		return
	}

	pw := user.Password

	pmc := &models.Pmc{
		Name: user.PmcName,
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

	tokenRes, err := authenticate(user.Email, pw)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&SignUpResponse{
			Error: err.Error(),
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
