package merchants

import (
	"github.com/gorilla/mux"
	"kiplingkelvin/golang-skeleton/internal/config"
	"kiplingkelvin/golang-skeleton/internal/merchants/handlers"
	"net/http"
)

// Payload ...
type Payload struct {
	Router *mux.Router
	Config *config.WebServerConfig
}

// InitializeRoutes ...
func InitializeRoute(payload Payload) {

	//handlers
	payload.Router.HandleFunc("/registration", handlers.MerchantRegistrationHandler).Methods(http.MethodPost)
	payload.Router.HandleFunc("/profile", handlers.ProfileUpdateHandler).Methods(http.MethodPut)
	payload.Router.HandleFunc("/profile", handlers.ProfileGetHandler).Methods(http.MethodGet)

}
