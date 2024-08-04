package main

import (
	"net/http"

	"github.com/eminsonlu/go-blog-aggregator/internal/auth"
	"github.com/eminsonlu/go-blog-aggregator/internal/database"
)

func (cfg *apiConfig) middlewareAuth(handler func(w http.ResponseWriter, r *http.Request, u database.User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		user, err := cfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		handler(w, r, user)
	}
}
