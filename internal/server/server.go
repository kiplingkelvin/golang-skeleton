package server

import (
	"fmt"
	"kiplingkelvin/golang-skeleton/internal/config"
	"kiplingkelvin/golang-skeleton/internal/models"
	"kiplingkelvin/golang-skeleton/internal/pkg"
	"kiplingkelvin/golang-skeleton/internal/pkg/postgres"

	"net/http"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

// This time make models.BookModel the dependency in Env.
var AllModelDaos models.PostgresDAO

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

	err = pkg.Initialize(webServerConfig.Service)
	if err != nil {
		logrus.WithField("Error", err).Error("Error initializing service")
		return err
	}

	postgresDB, err := postgres.NewPostgres(&webServerConfig.PostgresConfig).Db()

	// Initalise all the model daos
	AllModelDaos.Healthcheck = models.NewHealthCheckModel(postgresDB)
	AllModelDaos.Merchant = models.NewMerchantModel(postgresDB)
	AllModelDaos.BrandAlias = models.NewBrandingModel(postgresDB)

	server := NewServer(webServerConfig)
	server.Router.InitializeRoutes(webServerConfig)

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"tenant", "*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "UPDATE", "OPTIONS", "DELETE", "PATCH"},
	})

	var handler http.Handler

	if webServerConfig.CorsEnabled {
		handler = c.Handler(*&server.Router.Router)
	} else {
		handler = *&server.Router.Router
	}

	if err := http.ListenAndServe(
		fmt.Sprintf("%v:%v", "", webServerConfig.Port),
		handler,
	); err != nil {
		panic(err)
	}

	return nil
}
