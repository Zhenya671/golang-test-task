package v1

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (api ApiV1) client(r *mux.Router) *mux.Router {
	r.HandleFunc("/sign-in", api.h.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/sign-up", api.h.SignUp).Methods(http.MethodPost)

	v1 := r.PathPrefix("/user").Subrouter()
	v1.Use(api.h.Middleware)
	v1.HandleFunc("/debt", api.h.PayOff).Methods(http.MethodPut)

	v2 := v1.PathPrefix("/task").Subrouter()
	v2.HandleFunc("/algo/{AlgoName}", api.h.SolveAlgo).Methods(http.MethodPost)

	return r
}
