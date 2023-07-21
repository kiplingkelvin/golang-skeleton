package handlers

import (
	"encoding/json"
	"kiplingkelvin/golang-skeleton/internal/models"
	"kiplingkelvin/golang-skeleton/internal/services"
	"net/http"

	"github.com/sirupsen/logrus"
)

// TestHandler ...
func TestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	err := services.Service.PostgresDAO.PostgresPing(r.Context())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ping error": err,
		}).Error("json encode error")
	}


	healthCheckResponse := models.HealthCheck{
		Status:    "pass",
		Version:   "1",
		ReleaseID: "0.0.1",
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(healthCheckResponse)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("json encode error")
	}
}