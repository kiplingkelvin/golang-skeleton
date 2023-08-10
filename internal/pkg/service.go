package pkg

import (
	"github.com/sirupsen/logrus"
)

var Service GolangSkeletonService

// GolangSkeletonService ...
type GolangSkeletonService struct {
	config *ServiceConfig
}

// ServiceConfig ...
type ServiceConfig struct {
}

// Initialize ...
func Initialize(c *ServiceConfig) (err error) {

	logrus.Info("Initializing Service")

	Service = GolangSkeletonService{
		config: c,
	}

	return nil
}
