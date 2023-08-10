package healthcheck

import (
	"kiplingkelvin/golang-skeleton/internal/config"
	"kiplingkelvin/golang-skeleton/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

// Router ...
type TestRouter struct {
	Router *mux.Router
	DAO    models.PostgresDAO
	Config *config.WebServerConfig
}

// InitializeRoutes ...
func InitializeRoute(mainRouter TestRouter) {
	//handlers
	mainRouter.Router.HandleFunc("/healthcheck", mainRouter.TestHandler).Methods(http.MethodGet)

}
