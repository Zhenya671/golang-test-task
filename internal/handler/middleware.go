package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (h Handler) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		token = strings.TrimSpace(token)
		if token == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}

		userID, login, err := h.middleware.ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("token is invalid")
			return
		}

		r.Header.Set("UserID", userID)
		r.Header.Set("UserName", login)
		next.ServeHTTP(w, r)
	})
}
