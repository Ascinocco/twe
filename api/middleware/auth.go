package middleware

import (
	"TheWarEconomy/api/utils"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const (
	Authorization string = "Authorization"
)

type UnautherizedResponse struct {
	Error string `json:"error"`
}

type Token struct {
	UserId string
	jwt.StandardClaims
}

func sendUnauthorizedResponse(w http.ResponseWriter) {
	response := UnautherizedResponse{Error: "Unauthorized"}
	w.WriteHeader(http.StatusForbidden)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func JwtVerification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		publicEndpoint := []string{"/user/create", "/session/create"}
		reqPath := r.URL.Path

		// if request in to public endpoint, pass through
		for _, value := range publicEndpoint {
			if value == reqPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get(Authorization)

		// no token
		if tokenHeader == "" {
			sendUnauthorizedResponse(w)
			return
		}

		tokenParts := strings.Split(tokenHeader, " ")

		// malformed token, missing Bearer key or token value, or too many values...
		if len(tokenParts) != 2 {
			sendUnauthorizedResponse(w)
			return
		}

		clientToken := tokenParts[1]
		token := Token{}

		parsedToked, err := jwt.ParseWithClaims(clientToken, token, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv(utils.EnvTokenSecret)), nil
		})

		if err != nil {
			sendUnauthorizedResponse(w)
			return
		}

		if !parsedToked.Valid {
			sendUnauthorizedResponse(w)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", token.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
