package postgres

import (
	"context"

	"gorm.io/gorm"
)

type DAO interface {
    Connect() (*gorm.DB, error)
    Ping() error
	Db() (*gorm.DB, error)

    PostgresPing(ctx context.Context) error
}