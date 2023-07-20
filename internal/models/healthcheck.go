package models
type HealthCheck struct {
	Status    string   `json:"status"`
	Version   string   `json:"version"`
	ReleaseID string   `json:"release_id"`
}
