package postgres

import "context"

type DataAccess interface {
	Create(ctx context.Context, condition interface{}, model interface{}) (interface{}, error)
	Update(ctx context.Context, condition interface{}, model interface{}) error
	Get(ctx context.Context, condition interface{}) (interface{}, error)
	GetAll(ctx context.Context, model interface{}) ([]interface{}, error)
}
