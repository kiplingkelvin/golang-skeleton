package healthcheck

import (
	"encoding/json"
	"kiplingkelvin/golang-skeleton/internal/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

// TestHandler ...
func (ts *TestRouter) TestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	healthCheckResponse := models.HealthCheck{
		Status:    "pass",
		Version:   "2",
		ReleaseID: "0.0.2",
	}

	ts.DAO.Healthcheck.Create(r.Context(), healthCheckResponse)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(healthCheckResponse)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("json encode error")
	}
}
