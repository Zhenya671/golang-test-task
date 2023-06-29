package handler

import (
	"encoding/json"
	"github.com/Zhenya671/golang-test-task/internal/model"
	"net/http"
)

func (h Handler) handle_errors(w http.ResponseWriter, statusCode int, err string) {
	var newError model.ErrorResponse
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	newError.StatusCode = statusCode
	newError.Error = err
	marshal, err2 := json.Marshal(newError)

	if err2 != nil {
		return
	}

	w.Write(marshal)
}
