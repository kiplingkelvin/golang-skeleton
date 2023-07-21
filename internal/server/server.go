package server

import (
	"fmt"
	"kiplingkelvin/golang-skeleton/internal/config"
	"kiplingkelvin/golang-skeleton/internal/services"
	"net/http"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	Configuration *config.WebServerConfig
	Router        *Router
}

// NewServer ...
func NewServer(config *config.WebServerConfig) *Server {
	server := &Server{
		Configuration: config,
		Router:        NewRouter(),
	}

	return server
}

// RunServer ...
func RunServer() (err error) {
	webServerConfig, err := config.FromEnv()
	if err != nil {
		return err
	}

	logrus.Infof("Starting HTTPS server on port %s", webServerConfig.Port)

	err = services.Initialize(webServerConfig.Service)
	if err != nil {
		logrus.WithField("Error", err).Error("Error initializing service")
		return err
	}

	server := NewServer(webServerConfig)
	server.Router.InitializeRoutes(webServerConfig)

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"tenant", "*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "UPDATE", "OPTIONS", "DELETE", "PATCH"},
	})

	var handler http.Handler

	if webServerConfig.CorsEnabled {
		handler = c.Handler(*server.Router)
	} else {
		handler = *server.Router
	}

	if err := http.ListenAndServe(
		fmt.Sprintf("%v:%v", "", webServerConfig.Port),
		handler,
	); err != nil {
		panic(err)
	}

	return nil
}
