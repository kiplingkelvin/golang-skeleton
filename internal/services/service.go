package services

import (
    "github.com/sirupsen/logrus"
	postgresService "kiplingkelvin/golang-skeleton/internal/services/postgres"
)

var Service GolangSkeletonService

// GolangSkeletonService ...
type GolangSkeletonService struct {
	config      *ServiceConfig
    PostgresDAO postgresService.DAO
}

// ServiceConfig ...
type ServiceConfig struct {
    PostgresConfig *postgresService.Config `required:"true" split_words:"true"`
}

// Initialize ...
func Initialize(c *ServiceConfig) (err error) {

    logrus.Info("Initializing Service")

	Service = GolangSkeletonService{
		config: c,
	}

    Service.PostgresDAO = postgresService.NewPostgres(Service.config.PostgresConfig)

    return nil
}