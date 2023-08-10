package models

import "context"

type PostgresDAO struct {
	Healthcheck interface {
		Create(ctx context.Context, healthcheck HealthCheck) (*string, error)
		GetAll(context.Context) (*[]HealthCheck, error)
	}

	Merchant interface {
		Create(context.Context, Merchant) (*uint, error)
		Update(context.Context, Merchant) error
		Get(context.Context, Merchant) (*Merchant, error)
		GetAll(context.Context) (*[]Merchant, error)
	}

	BrandAlias interface {
		Create(context.Context, Branding) (*uint, error)
		Update(context.Context, Branding) error
		Get(context.Context, Branding) (*Branding, error)
		GetAll(context.Context) (*[]Branding, error)
	}
}
