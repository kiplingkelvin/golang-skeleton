package merchants

import (
	"kiplingkelvin/golang-skeleton/internal/config"
	"net/http"

	"github.com/gorilla/mux"
)

// Payload ...
type Payload struct {
	Router *mux.Router
	Config *config.WebServerConfig
}

// InitializeRoutes ...
func InitializeRoute(payload Payload) {

	//handlers
	payload.Router.HandleFunc("/registration", MerchantRegistrationHandler).Methods(http.MethodPost)
	payload.Router.HandleFunc("/profile", ProfileUpdateHandler).Methods(http.MethodPut)
	payload.Router.HandleFunc("/profile", ProfileGetHandler).Methods(http.MethodGet)

}
