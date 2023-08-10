package healthcheck

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"kiplingkelvin/golang-skeleton/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type HealthCheckModel struct {
	DB *gorm.DB
}

func NewHealthCheckModel(db *gorm.DB) *HealthCheckModel {
	return &HealthCheckModel{
		DB: db,
	}
}

func (dao *HealthCheckModel) Create(ctx context.Context, healthcheck models.HealthCheck) (*string, error) {
	releaseID := "1"
	return &releaseID, nil
}

func (dao *HealthCheckModel) GetAll(ctx context.Context) (*[]models.HealthCheck, error) {
	var healthcheck []models.HealthCheck

	healthcheck = append(healthcheck, models.HealthCheck{"pass", "1.2", "1"})
	healthcheck = append(healthcheck, models.HealthCheck{"pass", "2.2", "2"})

	return &healthcheck, nil
}

func TestHealthCheckHandler_whenRequestIsValid_ReturnsValidResponse(t *testing.T) {

	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	testWriter := new(httptest.ResponseRecorder)
	testResponseBodyBuffer := new(bytes.Buffer)
	testWriter.Body = testResponseBodyBuffer

	url := "/v1/healthcheck"

	testRequest, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		fmt.Println("bad resquest build")
	}

	testRequest.Header.Set("Accept", "application/json")
	testRequest.Header.Set("Accept-Encoding", "gzip,deflate,br")
	testRequest.Header.Set("Connection", "keep-alive")
	testRequest.Header.Set("Cache-Control", "no-cache")
	testRequest.Header.Set("Content-Type", "application/json")

	var router TestRouter
	router.DAO = models.PostgresDAO{
		Healthcheck: NewHealthCheckModel(db),
	}

	router.TestHandler(testWriter, testRequest)

	testResponse := &models.HealthCheck{}
	err = json.Unmarshal(testWriter.Body.Bytes(), testResponse)
	assert.Nil(t, err, "HealthCheckHandler: Expected error to be nil when unmarshaling")

	assert.Equal(t, "pass", testResponse.Status)
	assert.Equal(t, "2", testResponse.Version)
	assert.Equal(t, "0.0.2", testResponse.ReleaseID)
}
