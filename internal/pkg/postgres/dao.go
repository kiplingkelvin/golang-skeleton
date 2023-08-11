package postgres

import "context"


type DataAccess interface {
	Create(context.Context, interface{}) (*uint, error)
	Update(context.Context, interface{}) error
	Get(context.Context, interface{}) (*interface{}, error)
	GetAll(context.Context) (*[]interface{}, error)
}

