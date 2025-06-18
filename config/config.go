package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	SQL    SQL    `yaml:"sql"`
	Server Server `yaml:"server"`
}

type Server struct {
	Port            int    `yaml:"port"`
	Host            string `yaml:"host"`
	AuthTokenSecret string `yaml:"auth_token_secret"`
}

type SQL struct {
	PrimaryDatabase SqlSettings `yaml:"primary_database"`
}

type SqlSettings struct {
	Host                  string  `yaml:"host"`
	Port                  int     `yaml:"port"`
	DbName                string  `yaml:"db_name"`
	UserName              string  `yaml:"username"`
	Password              string  `yaml:"password"`
	MaxOpenConnections    int     `yaml:"max_open_connections"`
	MaxIdleConnections    int     `yaml:"max_idle_connections"`
	MaxIdleConnectionTime float64 `yaml:"max_idle_connection_time"`
}

func New() (conf Config, _ error) {
	loadConfig := func(fileName string) error {
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(fileData, &conf)
		if err != nil {
			return err
		}
		return nil
	}

	if err := loadConfig("config.yaml"); err != nil {
		return conf, fmt.Errorf("reading config file: %w", err)
	}
	bindFromEnv(&conf)
	return
}
func bindFromEnv(conf *Config) {
	setIfNotSet := func(field *string, envName string) {
		if os.Getenv(envName) != "" {
			*field = os.Getenv(envName)
		}
	}
	setIfNotSet(&conf.Server.AuthTokenSecret, "AUTH_TOKEN_SECRET")
	setIfNotSet(&conf.SQL.PrimaryDatabase.Password, "PRIMARY_DB_PASSWORD")
}
