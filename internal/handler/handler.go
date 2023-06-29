package handler

import (
	"github.com/Zhenya671/golang-test-task/internal/service"
	"github.com/Zhenya671/golang-test-task/internal/settings"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	user       service.IUserService
	middleware service.IMiddleWere
	log        *logrus.Logger
}

func NewHandler(userService service.IUserService, log *logrus.Logger, settings settings.AppSettings) *Handler {
	return &Handler{
		user:       userService,
		log:        log,
		middleware: service.UserService{AppSettings: settings},
	}
}
