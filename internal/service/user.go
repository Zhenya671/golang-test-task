package service

import (
	"github.com/Zhenya671/golang-test-task/internal/messages"
	"github.com/Zhenya671/golang-test-task/internal/model"
	"github.com/Zhenya671/golang-test-task/internal/repository"
	"github.com/Zhenya671/golang-test-task/internal/settings"
	"github.com/sirupsen/logrus"
	"strings"
)

type IUserService interface {
	SignIn(logIn model.User) (string, error)
	SignUp(user model.User) (string, error)
	PayOff(userId string, input model.Debt) (model.Debt, error)
}

type UserService struct {
	repo   repository.IUserRepository
	logger *logrus.Logger
	settings.AppSettings
}

func NewUserService(repo repository.IUserRepository, logger *logrus.Logger, appSettings settings.AppSettings) *UserService {
	return &UserService{repo: repo, logger: logger, AppSettings: appSettings}
}

func (s UserService) SignIn(logIn model.User) (string, error) {
	password := s.generatePasswordHash(logIn.Password)
	logIn.Password = password
	user, err := s.repo.SignIn(logIn)
	if err != nil {
		s.logger.Warn(err)
		return "", messages.AppErrorUserNotFound
	}

	token, err := s.GenerateToken(user.ID, user.Login)
	if err != nil {
		s.logger.Warn(err)
		return "", messages.AppErrorStatusBadRequest
	}

	return token, nil
}

func (s UserService) SignUp(newUser model.User) (string, error) {
	password := s.generatePasswordHash(newUser.Password)
	newUser.Password = password
	newUser.Login = strings.ToUpper(newUser.LastName) + strings.ToLower(newUser.FirstName)

	user, err := s.repo.SignUp(newUser)
	if err != nil {
		s.logger.Warn(err)
		return "", messages.AppErrorSuchUserExist

	}

	token, err := s.GenerateToken(user.ID, user.Login)
	if err != nil {
		s.logger.Warn(err)
		return "", messages.AppErrorStatusBadRequest
	}

	return token, nil
}

func (s UserService) PayOff(userId string, input model.Debt) (model.Debt, error) {
	var debt model.Debt

	debt, err := s.repo.PayOff(userId, input)
	if err != nil {
		s.logger.Warn(err)
		return debt, err
	}

	return debt, nil
}
