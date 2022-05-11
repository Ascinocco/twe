package controllers

import (
	"TheWarEconomy/api/database"
	"TheWarEconomy/api/middleware"
	"TheWarEconomy/api/models"
	"TheWarEconomy/api/utils"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type SessionResponse struct {
	Token string `json:"token"`
	Error string `json:"error"`
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
		return "", errors.New("Unable to login, please try again")
	}

	if err != nil {
		return "", errors.New("Unable to login, please try again")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("Unable to login, please try again")
	}

	tokenData := &middleware.Token{UserId: user.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenData)
	tokenString, err := token.SignedString([]byte(os.Getenv(utils.EnvTokenSecret)))

	if err != nil {
		return "", errors.New("Unable to login, please try again")
	}

	return tokenString, nil
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&SessionResponse{
			Error: "Unable to login, please try again",
		})
		return
	}

	tokenRes, err := authenticate(user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&SessionResponse{
			Error: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SessionResponse{
		Token: tokenRes,
		Error: "",
	})
}
