package utils

import (
	"errors"
	"net/http"
)

func GetUserIdFromCtx(r *http.Request) (string, error) {
	ctxUserId := r.Context().Value("userId")

	switch v := ctxUserId.(type) {
	case string:
		return v, nil
	default:
		return "", errors.New("Unable to authorize request")
	}
}
