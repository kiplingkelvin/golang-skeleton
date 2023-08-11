package postgres

import (
	"context"
	"errors"
	"fmt"
	"kiplingkelvin/golang-skeleton/internal/merchants/models"
	"reflect"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Service PostgresService

type PostgresService struct {
	DAO DataAccess
}

var pg Postgres

type Postgres struct {
	db     *gorm.DB
	config *Config
}

type Config struct {
	DatabaseName string `envconfig:"DB_DATABASE_NAME" required:"true" split_words:"true" default:"lordofrings"`
	User         string `envconfig:"DB_DATABASE_USER" required:"true" split_words:"true" default:"kuria_kdb"`
	Password     string `envconfig:"DB_DATABASE_PASSWORD" required:"true" split_words:"true" default:"ThisIsAVeryStrongPassword"`
	Host         string `envconfig:"DB_DATABASE_HOST" required:"true" split_words:"true" default:"chpter-db-jarvis.c9iafjhtw7p1.us-east-1.rds.amazonaws.com"`
	Port         uint32 `envconfig:"DB_DATABASE_PORT" required:"false" split_words:"true" default:"5432" `
}

func InitDB(config *Config) error {
	pg.config = config

	//Open DB connection
	db, err := Connect()
	if err != nil {
		return err
	}
	pg.db = db
	db.AutoMigrate(&models.Merchant{})

	//Initialize the DAOs
	Service.DAO = &pg
	return nil
}

func Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(getConnectionString()), nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", pg.config.Host, pg.config.User, pg.config.Password, pg.config.DatabaseName, pg.config.Port)
}

func Ping() error {
	db, err := Connect()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func Migrations() error {
	if pg.db == nil {
		db, err := Connect()
		if err != nil {
			return err
		}

		//If connection is okay Run Migrations
		db.AutoMigrate(&models.Merchant{})
	}
	return nil
}

func (dao *Postgres) Create(ctx context.Context, condition interface{}, model interface{}) (interface{}, error) {
	// Use reflection to create a new pointer to the type of model
	modelPtr := reflect.New(reflect.TypeOf(model))

	// Dereference the pointer and set the value
	modelPtr.Elem().Set(reflect.ValueOf(model))

	tx := dao.db.Where(condition).FirstOrCreate(modelPtr.Interface())

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("exists")
	}

	return nil, nil
}

func (dao *Postgres) Update(ctx context.Context, condition interface{}, model interface{}) error {
	// Use reflection to create a new pointer to the type of model
	modelPtr := reflect.New(reflect.TypeOf(model))

	// Dereference the pointer and set the value
	modelPtr.Elem().Set(reflect.ValueOf(model))

	tx := dao.db.Model(&model).Where(condition).Updates(modelPtr.Interface())
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (dao *Postgres) Get(ctx context.Context, condition interface{}) (interface{}, error) {
	// Use reflection to create a new pointer to the type of model
	modelPtr := reflect.New(reflect.TypeOf(condition))

	// Call the database query using the pointer to the model
	tx := dao.db.Where(condition).First(modelPtr.Interface())

	if tx.Error != nil {
		return nil, tx.Error
	}

	return modelPtr.Elem().Interface(), nil
}

func (dao *Postgres) GetAll(ctx context.Context, model interface{}) ([]interface{}, error) {
	// Create a slice to store the results
	results := make([]interface{}, 0)

	// Use reflection to create a new pointer to a slice of the model type
	slicePtr := reflect.New(reflect.SliceOf(reflect.TypeOf(model)))

	// Call the database query to retrieve all records
	tx := dao.db.Find(slicePtr.Interface())

	if tx.Error != nil {
		return nil, tx.Error
	}

	// Convert the slice pointer to a slice value
	sliceValue := slicePtr.Elem()

	// Iterate through the slice and append each element to the results
	for i := 0; i < sliceValue.Len(); i++ {
		results = append(results, sliceValue.Index(i).Addr().Interface())
	}

	return results, nil
}
