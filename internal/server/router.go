package server

import (
	"kiplingkelvin/golang-skeleton/internal/config"
	"kiplingkelvin/golang-skeleton/internal/merchants"
	"kiplingkelvin/golang-skeleton/internal/pkg/postgres"

	"github.com/gorilla/mux"
)

// Router ...
type Router struct {
	Router *mux.Router
	DAO    postgres.PostgresDAO
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

	merchants.InitializeRoute(merchants.Payload{Router: route, DAO: r.DAO, Config: cfg})
}
