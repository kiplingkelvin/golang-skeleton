package server

import (
	"github.com/gorilla/mux"
	"kiplingkelvin/golang-skeleton/internal/config"
	"kiplingkelvin/golang-skeleton/internal/merchants"
)

// Router ...
type Router struct {
	Router *mux.Router
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{
		Router: mux.NewRouter(),
	}
}

// InitializeRoutes ...
func (r *Router) InitializeRoutes(cfg *config.WebServerConfig) {
	route := r.Router.PathPrefix("/v1").Subrouter()

	merchants.InitializeRoute(merchants.Payload{Router: route, Config: cfg})
}
