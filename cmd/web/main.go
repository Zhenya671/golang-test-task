package main

import (
	"context"
	"fmt"
	v1 "github.com/Zhenya671/golang-test-task/internal/api/v1"
	"github.com/Zhenya671/golang-test-task/internal/handler"
	"github.com/Zhenya671/golang-test-task/internal/repository"
	"github.com/Zhenya671/golang-test-task/internal/service"
	"github.com/Zhenya671/golang-test-task/internal/settings"
	"github.com/Zhenya671/golang-test-task/internal/usecases"
	"github.com/Zhenya671/golang-test-task/migration"
	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	migrationPath = "/etc/migration"
)

type server struct {
	server *http.Server
}

func (s *server) swaggerRun(host string) {
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://213.139.210.171:12/swagger/doc.json"), //The url pointing to API definition
	))

	fmt.Println(http.ListenAndServe(":12", r))

}

func (s *server) run(r *mux.Router, conf *settings.AppSettings) error {
	handlerCors := cors.AllowAll().Handler(r)
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.Port),
		Handler: handlerCors,
	}

	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *server) shutdown(log *logrus.Logger) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	err := s.server.Shutdown(ctx)
	if err != nil {
		log.Errorf("can't shutdown server: \n%s", err.Error())
	}
	defer cancel()
}

// @title Your Application API
// @description This is the API documentation for Your Application.
// @version 1.0
// @host localhost:8080
// @BasePath /

func main() {
	logger := logrus.New()
	logger.Info("App starting")

	client := http.Client{}
	_ = client

	newAppSettings, err := settings.NewAppSettings()
	if err != nil {
		logger.Fatalf("can't to get settings:  %s", err.Error())
	}
	logger.Infof("App configs: %s", newAppSettings)

	newRepository, err := repository.NewRepository(newAppSettings.PgConf)
	if err != nil {
		logger.Fatalf("can't connect to db: %s", err)
	}
	logger.Info(*newRepository)

	if err := migration.NewMigration(newRepository.DB, logger, migrationPath); err != nil {
		logger.Fatalf("can't proceed migration: %s", err)
	}

	newUserService := service.NewUserService(newRepository, logger, *newAppSettings)
	newHandler := handler.NewHandler(newUserService, logger, *newAppSettings)
	router := v1.NewApiV1(newHandler, logger)

	srv := new(server)

	go srv.shutdown(logger)

	ip := usecases.GetOutboundIP().String()
	go srv.swaggerRun(ip)

	logger.Infof("swagger started url %s", fmt.Sprintf("http://localhost:12/swagger/index.html", ip))
	logger.Infof("application started on %s:%s", ip, newAppSettings.Port)

	if err := srv.run(router, newAppSettings); err != nil {
		logger.Fatalf("can't serve server: \n%s", err.Error())
	}
}
