package handlers

import (
	"net/http"
    "encoding/json"
	"github.com/sirupsen/logrus"
    "kiplingkelvin/golang-skeleton/internal/models"
)

// TestHandler ...
func TestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	healthCheckResponse := models.HealthCheck{
		Status:    "pass",
		Version:   "1",
		ReleaseID: "0.0.1",
	}

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(healthCheckResponse)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("json encode error")
	}
}