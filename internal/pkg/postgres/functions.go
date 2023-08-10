package postgres

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// Create a custom MerchantModel type which wraps the gorm.DB connection pool.
type DbInit struct {
	DB *gorm.DB
}

func NewDbInit(db *gorm.DB) *DbInit {
	return &DbInit{
		DB: db,
	}
}

func (dao *DbInit) Create(ctx context.Context, model interface{}) (*uint, error) {

	tx := dao.DB.FirstOrCreate(&model)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("exists")
	}

	return nil, nil
}

func (dao *DbInit) Update(ctx context.Context, model interface{}) error {
	tx := dao.DB.Model(&model).Updates(model)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (dao *DbInit) Get(ctx context.Context, model interface{}) (*interface{}, error) {
	var data *interface{}
	tx := dao.DB.Model(model).First(&data)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return data, nil
}

func (dao *DbInit) GetAll(ctx context.Context) (*[]interface{}, error) {
	var data []interface{}
	tx := dao.DB.Find(&data)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("not found")
	}

	return &data, nil
}
