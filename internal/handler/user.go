package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Zhenya671/golang-test-task/internal/messages"
	"github.com/Zhenya671/golang-test-task/internal/model"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strings"
)

// SignUp handles the user sign-up request.
// @Summary User sign-up
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param input body model.User true "User registration data"
// @Success 200 {object} model.Token "Token response"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Router /sign-up [post]
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

// SignIn handles the user sign-in request.
// @Summary User sign-in
// @Description Authenticate user credentials and generate a token
// @Tags User
// @Accept json
// @Produce json
// @Param input body model.User true "User credentials"
// @Success 200 {object} model.Token "Token response"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Router /sign-in [post]
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

// PayOff handles the debt payment request.
// @Summary Pay off debt
// @Description Process debt payment for a user
// @Tags User
// @Accept json
// @Produce json
// @Param UserID header string true "User ID"
// @Param input body model.Debt true "Debt details"
// @Success 200 {object} model.Debt "Payment response"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Router /user/pay-off [post]
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

// SolveAlgo godoc
// @Summary Solve Algorithm
// @Description Solve an algorithm task
// @Tags Algo
// @Accept json
// @Produce json
// @Param UserID header string true "User ID"
// @Param AlgoName path string true "Algorithm Name"
// @Param requestTask body model.Task true "Request Body"
// @Success 200 {object} model.Task
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/task/algo/{AlgoName} [post]
func (h Handler) SolveAlgo(w http.ResponseWriter, r *http.Request) {
	var requestTask model.Task

	userID := r.Header.Get("UserID")
	if userID == "" {
		h.handle_errors(w, http.StatusUnauthorized, messages.AppErrorUnauthorized.Error())
		return
	}

	algoName := mux.Vars(r)["AlgoName"]
	if algoName == "" {
		h.handle_errors(w, http.StatusBadRequest, messages.AppErrorStatusBadRequest.Error())
		return
	}
	algoName = strings.ReplaceAll(algoName, "-", " ")
	algoName = strings.Title(algoName)
	algoName = strings.ReplaceAll(algoName, " ", "")
	fmt.Println(algoName)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Warn("Failed to read request body:", err)
		h.handle_errors(w, http.StatusBadRequest, messages.AppErrorStatusBadRequest.Error())
		return
	}

	err = json.Unmarshal(body, &requestTask)
	if err != nil {
		log.Println("Failed to unmarshal request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.user.SolveAlgo(userID, algoName, requestTask)
	if err != nil {
		h.log.Warn(err)
		h.handle_errors(w, http.StatusBadRequest, messages.AppErrorCantSolveTask.Error())
		return
	}

	marshal, err := json.Marshal(result)
	if err != nil {
		h.handle_errors(w, http.StatusBadRequest, messages.AppErrorWithMarshalling.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}
