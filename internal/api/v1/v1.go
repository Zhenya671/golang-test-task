package v1

import (
	"net/http"

	"github.com/Zhenya671/golang-test-task/internal/handler"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiV1 struct {
	h      *handler.Handler
	logger *logrus.Logger
}

func NewApiV1(handler *handler.Handler, logger *logrus.Logger) *mux.Router {
	v1 := ApiV1{h: handler, logger: logger}
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	rV1 := api.PathPrefix("/v1").Subrouter()
	rV1.Use(v1.reqLogsFunc)

	v1.client(rV1)

	return r
}

func (api *ApiV1) reqLogsFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.logger.Info(
			"request information: method ->", r.Method,
			", url ->", r.URL,
			", token ->", r.Header.Get("token"))
		next.ServeHTTP(w, r)
	})
}
