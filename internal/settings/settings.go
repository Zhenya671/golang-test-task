package settings

import (
	"fmt"
	"os"
)

type AppSettings struct {
	Port     string
	HostName string
	TokenKey string
	PgConf   string
}

type PostgresConfig struct {
	Port     string
	Host     string
	Name     string
	Username string
	Password string
	SSLMode  string
}

func NewAppSettings() (*AppSettings, error) {
	var settings AppSettings
	var pgConfig PostgresConfig
	var err error

	settings.Port, err = getEnvironmentVariables("APP_PORT")
	if err != nil {
		return nil, err
	}
	settings.HostName, err = getEnvironmentVariables("APP_HOSTNAME")
	if err != nil {
		return nil, err
	}
	settings.HostName, err = getEnvironmentVariables("TOKEN_KEY")
	if err != nil {
		return nil, err
	}
	pgConfig.Host, err = getEnvironmentVariables("DB_HOSTNAME")
	if err != nil {
		return nil, err
	}
	pgConfig.Port, err = getEnvironmentVariables("DB_PORT")
	if err != nil {
		return nil, err
	}
	pgConfig.Name, err = getEnvironmentVariables("DB_NAME")
	if err != nil {
		return nil, err
	}
	pgConfig.Username, err = getEnvironmentVariables("DB_USERNAME")
	if err != nil {
		return nil, err
	}
	pgConfig.Password, err = getEnvironmentVariables("DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	pgConfig.SSLMode, err = getEnvironmentVariables("DB_SSLMODE")
	if err != nil {
		return nil, err
	}

	settings.PgConf = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%s",
		pgConfig.Username, pgConfig.Password, pgConfig.Host, pgConfig.Port,
		pgConfig.Name, pgConfig.SSLMode)

	return &settings, nil
}

func getEnvironmentVariables(param string) (string, error) {
	variable := os.Getenv(param)
	if variable == "" {
		return "", fmt.Errorf("no environment variables with name %s", param)
	}
	return variable, nil

}
