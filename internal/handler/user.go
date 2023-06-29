package handler

import (
	"encoding/json"
	"github.com/Zhenya671/golang-test-task/internal/messages"
	"github.com/Zhenya671/golang-test-task/internal/model"
	"io"
	"net/http"
)

func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input model.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.log.Warn(err)
		h.handle_errors(w, http.StatusBadRequest, "Bad Request")
		return
	}

	token, err := h.user.SignUp(input)
	if err != nil {
		h.log.Warn(err)
		h.handle_errors(w, http.StatusBadRequest, err.Error())
		return
	}
	var tokenJSON model.Token
	tokenJSON.Token = token

	marshal, err := json.Marshal(tokenJSON)
	if err != nil {
		h.log.Warn(err)
		h.handle_errors(w, http.StatusBadRequest, messages.AppErrorWithMarshalling.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(marshal)
}

func (h Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input model.User

	all, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Warn(err)
		h.handle_errors(w, http.StatusBadRequest, messages.AppErrorWithMarshalling.Error())
		return
	}

	err = json.Unmarshal(all, &input)
	if err != nil {
		h.log.Warn(err)
		h.handle_errors(w, http.StatusBadRequest, "Bad Request")
		return
	}

	token, err := h.user.SignIn(input)
	if err != nil {
		h.log.Warn(err)
		h.handle_errors(w, http.StatusBadRequest, err.Error())
		return
	}

	var tokenJSON model.Token
	tokenJSON.Token = token

	marshal, err := json.Marshal(tokenJSON)
	if err != nil {
		h.log.Warn(err)
		h.handle_errors(w, http.StatusBadRequest, messages.AppErrorWithMarshalling.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(marshal)
}

func (h Handler) PayOff(w http.ResponseWriter, r *http.Request) {
	var input model.Debt

	userID := r.Header.Get("UserID")
	if userID == "" {
		h.handle_errors(w, http.StatusUnauthorized, messages.AppErrorUnauthorized.Error())
		return
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.log.Warn(err)
		h.handle_errors(w, http.StatusBadRequest, messages.AppErrorWithMarshalling.Error())
		return
	}

	payOff, err := h.user.PayOff(userID, input)
	if err != nil {
		h.handle_errors(w, http.StatusBadRequest, messages.AppErrorCantPayOff.Error())
		return
	}

	marshal, err := json.Marshal(payOff)
	if err != nil {
		h.handle_errors(w, http.StatusBadRequest, messages.AppErrorWithMarshalling.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}
