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
	payload.Router.HandleFunc("/registration", payload.MerchantRegistrationHandler).Methods(http.MethodPost)
	payload.Router.HandleFunc("/profile", payload.ProfileUpdateHandler).Methods(http.MethodPut)
	payload.Router.HandleFunc("/profile", payload.ProfileGetHandler).Methods(http.MethodGet)

}
