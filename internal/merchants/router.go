package merchants

import (
	"kiplingkelvin/golang-skeleton/internal/config"
	"kiplingkelvin/golang-skeleton/internal/pkg/postgres"
	"net/http"

	"github.com/gorilla/mux"
)

// Payload ...
type Payload struct {
	Router *mux.Router
	DAO    postgres.PostgresDAO
	Config *config.WebServerConfig
}

// InitializeRoutes ...
func InitializeRoute(payload Payload) {

	//handlers
	payload.Router.HandleFunc("/merchant-create", payload.MerchantRegistrationHandler).Methods(http.MethodPost)
	payload.Router.HandleFunc("/merchant-update", payload.ProfileUpdateHandler).Methods(http.MethodPut)
	payload.Router.HandleFunc("/merchant-get", payload.ProfileGetHandler).Methods(http.MethodGet)

}
