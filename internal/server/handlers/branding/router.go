package branding

import (
	"kiplingkelvin/golang-skeleton/internal/config"
	"kiplingkelvin/golang-skeleton/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

// Payload ...
type Payload struct {
	Router *mux.Router
	DAO    models.PostgresDAO
	Config *config.WebServerConfig
}

// InitializeRoutes ...
func InitializeRoute(payload Payload) {

	//handlers
	payload.Router.HandleFunc("/branding", payload.CreateBrandingHandler).Methods(http.MethodPost)
	payload.Router.HandleFunc("/branding", payload.UpdateBrandingHandler).Methods(http.MethodPut)
	payload.Router.HandleFunc("/branding", payload.GetBrandingHandler).Methods(http.MethodGet)

}
