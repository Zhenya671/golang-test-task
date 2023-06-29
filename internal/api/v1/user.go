package v1

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (api ApiV1) client(r *mux.Router) *mux.Router {
	r.HandleFunc("/sign-in", api.h.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/sign-up", api.h.SignUp).Methods(http.MethodPost)

	v1 := r.PathPrefix("").Subrouter()
	v1.Use(api.h.Middleware)

	return r
}
