package postgres

import (
	"context"
	"errors"
	"reflect"

	"github.com/sirupsen/logrus"
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

	// Start a new transaction
	tx := dao.DB.Begin()

	// Use reflection to create a new pointer to the type of model
	modelPtr := reflect.New(reflect.TypeOf(model))

	// Dereference the pointer and set the value
	modelPtr.Elem().Set(reflect.ValueOf(model))

	// Call FirstOrCreate with the pointer to the model
	result := tx.FirstOrCreate(modelPtr.Interface())

	if result.Error != nil {
		tx.Rollback() // Rollback transaction if there's an error
		return nil, result.Error
	}

	if result.RowsAffected != 1 {
		tx.Rollback() // Rollback transaction if the row does not exist
		return nil, errors.New("exists")
	}

	// Commit the transaction
	tx.Commit()

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
