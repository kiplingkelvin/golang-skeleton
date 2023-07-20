package server

import (
    "net/http"
    "github.com/gorilla/mux"
	"kiplingkelvin/golang-skeleton/internal/config"
	"kiplingkelvin/golang-skeleton/internal/server/handlers"
)

// Router ...
type Router struct {
	*mux.Router
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

// InitializeRoutes ...
func (r *Router) InitializeRoutes(cfg *config.WebServerConfig) {

	route := r.Router.PathPrefix("/v1").Subrouter()

	//handlers
	route.HandleFunc("/test", handlers.TestHandler).Methods(http.MethodGet)

}
