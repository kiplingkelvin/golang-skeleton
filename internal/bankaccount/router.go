package bankaccount

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
	payload.Router.HandleFunc("/bankaccount-create", payload.BankAccountCreateHandler).Methods(http.MethodPost)
	payload.Router.HandleFunc("/bankaccount-update", payload.BankAccountUpdateHandler).Methods(http.MethodPut)
	payload.Router.HandleFunc("/bankaccount-get", payload.BankAccountGetHandler).Methods(http.MethodGet)

}
