package models

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type HealthCheck struct {
	Status    string `json:"status"`
	Version   string `json:"version"`
	ReleaseID string `json:"release_id"`
}

// Create a custom HealthCheckModel type which wraps the gorm.DB connection pool.
type HealthCheckModel struct {
	DB *gorm.DB
}

func NewHealthCheckModel(db *gorm.DB) *HealthCheckModel {
	return &HealthCheckModel{
		DB: db,
	}
}

func (dao *HealthCheckModel) Create(ctx context.Context, healthcheck HealthCheck) (*string, error) {

	tx := dao.DB.Where("version = ?", healthcheck.Version).FirstOrCreate(&healthcheck)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("version exists")
	}

	return &healthcheck.ReleaseID, nil
}

func (dao *HealthCheckModel) GetAll(ctx context.Context) (*[]HealthCheck, error) {

	var healthcheck []HealthCheck
	tx := dao.DB.Find(&healthcheck)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("record not found")
	}

	return &healthcheck, nil
}
