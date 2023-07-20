package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

const (
	envPrefix = "GolangSkeleton" //TODO Get this from .env
)

// WebServerConfig ...
type WebServerConfig struct {
	Port        string `required:"true" split_words:"true"`
	EnableAuth  bool
	CorsEnabled bool   `default:"true" split_words:"true"`
	JWTSecret   string `required:"true" split_words:"true"`
}

// FromEnv ...
func FromEnv() (cfg *WebServerConfig, err error) {
	fromFileToEnv()

	cfg = &WebServerConfig{}

	err = envconfig.Process(envPrefix, cfg)
	if err != nil {

		return nil, err
	}

	return cfg, nil
}

func fromFileToEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warning("No config files found to load to env. Defaulting to environment.")
	}
}
