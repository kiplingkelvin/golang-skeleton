package postgres

import (
	"context"
	"errors"
	"fmt"
	merchants "kiplingkelvin/golang-skeleton/internal/merchants/models"
	bankaccount "kiplingkelvin/golang-skeleton/internal/bankaccount/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func NewPostgres(config *Config) *Postgres {
	return &Postgres{
		config: config,
	}
}


func (dao *Postgres) Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dao.getConnectionString()), nil)
	if err != nil {
		return nil, err
	}
	dao.db = db
	return db, nil
}

func (dao *Postgres) getConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dao.config.Host, dao.config.User, dao.config.Password, dao.config.DatabaseName, dao.config.Port)
}

func (dao *Postgres) Ping() error {
	db, err := dao.Connect()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func (dao *Postgres) Db() (*gorm.DB, error) {
	if dao.db == nil {
		db, err := dao.Connect()
		if err != nil {
			return nil, err
		}

		//If connection is okay Run Migrations
		db.AutoMigrate(&merchants.Merchant{})
		db.AutoMigrate(&bankaccount.BankAccount{})
		dao.db = db
	}
	return dao.db, nil
}



func (dao *Postgres) Create(ctx context.Context, model interface{}) (*uint, error) {

	tx := dao.db.FirstOrCreate(&model)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("exists")
	}

	return nil, nil
}

func (dao *Postgres) Update(ctx context.Context, model interface{}) error {
	tx := dao.db.Model(&model).Updates(model)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (dao *Postgres) Get(ctx context.Context, model interface{}) (*interface{}, error) {
	var data *interface{}
	tx := dao.db.Model(model).First(&data)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return data, nil
}

func (dao *Postgres) GetAll(ctx context.Context) (*[]interface{}, error) {
	var data []interface{}
	tx := dao.db.Find(&data)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("not found")
	}

	return &data, nil
}

