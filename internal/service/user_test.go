package service_test

import (
	"errors"
	"testing"

	"github.com/Zhenya671/golang-test-task/internal/model"
	mock "github.com/Zhenya671/golang-test-task/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	ErrInvalidCredentials       = errors.New("invalid credentials")
	ErrUserExists               = errors.New("user already exists")
	ErrInvalidUserID            = errors.New("invalid user ID")
	ErrInvalidUserIDOrAlgorithm = errors.New("invalid user ID or algorithm name")
)

func TestUserService_SignUp(t *testing.T) {
	userService := new(mock.IUserService)

	user := model.User{
		LastName:    "testlastname",
		FirstName:   "testfirstname",
		GroupNumber: "testgroupnumber",
		Login:       "testuser",
		Password:    "testpassword",
		Debt:        model.Debt{},
	}

	// Positive scenario: Valid user data
	userService.On("SignUp", user).Return("token", nil)
	token, err := userService.SignUp(user)

	userService.AssertExpectations(t)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Negative scenario: Existing user
	userService.On("SignUp", user).Return("", ErrUserExists)

	token1, err1 := "", ErrUserExists

	userService.AssertExpectations(t)
	assert.Error(t, err1)
	assert.Empty(t, token1)
}

func TestUserService_SignIn(t *testing.T) {
	userService := new(mock.IUserService)

	user := model.User{
		Login:    "testuser",
		Password: "testpassword",
	}

	t.Run("Positive scenario: Valid user credentials", func(t *testing.T) {
		userService.On("SignIn", user).Return("qweqwe", nil)

		token, err := userService.SignIn(user)

		userService.AssertExpectations(t)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("Negative scenario: Invalid user credentials", func(t *testing.T) {
		userService.On("SignIn", user).Return("", ErrInvalidCredentials)

		token, err := "", ErrInvalidCredentials

		userService.AssertExpectations(t)
		assert.Error(t, err)
		assert.Empty(t, token)
	})
}

func TestUserService_PayOff(t *testing.T) {
	userService := new(mock.IUserService)

	userID := "123"
	debt := model.Debt{Amount: 100.0}

	expectedDebt := model.Debt{Amount: 0.0}

	// Positive scenario: Valid user ID and debt input
	userService.On("PayOff", userID, debt).Return(expectedDebt, nil)

	resultDebt, err := userService.PayOff(userID, debt)

	userService.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedDebt, resultDebt)

	// Negative scenario: Invalid user ID
	userService.On("PayOff", userID, debt).Return(model.Debt{}, ErrInvalidUserID)

	resultDebt, err = model.Debt{}, ErrInvalidUserID

	userService.AssertExpectations(t)
	assert.Error(t, err)
	assert.Empty(t, resultDebt)
}

func TestUserService_SolveAlgo(t *testing.T) {
	userService := new(mock.IUserService)

	userID := "123"
	algorithmName := "testalgo"
	taskInput := model.Task{InputData: []interface{}{1, 1, 3, 3, 5, 6, 5, 5, 5, 4}}

	expectedTask := model.Task{InputData: []interface{}{2}}

	// Positive scenario: Valid user ID, algorithm name, and task input
	userService.On("SolveAlgo", userID, algorithmName, taskInput).Return(expectedTask, nil)

	resultTask, err := userService.SolveAlgo(userID, algorithmName, taskInput)

	userService.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedTask, resultTask)

	// Negative scenario: Invalid user ID or algorithm name
	userService.On("SolveAlgo", userID, algorithmName, taskInput).Return(model.Task{}, ErrInvalidUserIDOrAlgorithm)

	resultTask, err = model.Task{}, ErrInvalidUserIDOrAlgorithm

	userService.AssertExpectations(t)
	assert.Error(t, err)
	assert.Empty(t, resultTask)
}
