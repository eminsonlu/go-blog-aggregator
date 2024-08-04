package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrorNoAuth = errors.New("no auth header")

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrorNoAuth
	}

	splited := strings.Split(authHeader, " ")
	if len(splited) != 2 && splited[0] != "ApiKey" {
		return "", ErrorNoAuth
	}

	return splited[1], nil
}
