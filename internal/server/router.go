package server

import (
	"kiplingkelvin/golang-skeleton/internal/config"
	"kiplingkelvin/golang-skeleton/internal/models"
	"kiplingkelvin/golang-skeleton/internal/server/handlers/branding"
	"kiplingkelvin/golang-skeleton/internal/server/handlers/healthcheck"
	"kiplingkelvin/golang-skeleton/internal/server/handlers/merchants"

	"github.com/gorilla/mux"
)

// Router ...
type Router struct {
	Router *mux.Router
	DAO    models.PostgresDAO
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{
		Router: mux.NewRouter(),
		DAO:    AllModelDaos,
	}
}

// InitializeRoutes ...
func (r *Router) InitializeRoutes(cfg *config.WebServerConfig) {
	route := r.Router.PathPrefix("/v1").Subrouter()

	healthcheck.InitializeRoute(healthcheck.TestRouter{Router: route, DAO: r.DAO, Config: cfg})
	merchants.InitializeRoute(merchants.Payload{Router: route, DAO: r.DAO, Config: cfg})
	branding.InitializeRoute(branding.Payload{Router: route, DAO: r.DAO, Config: cfg})
}
