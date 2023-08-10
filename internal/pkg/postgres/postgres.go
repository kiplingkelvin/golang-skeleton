package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kiplingkelvin/golang-skeleton/internal/models"
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
		db.AutoMigrate(&models.HealthCheck{})
		dao.db = db
	}
	return dao.db, nil
}
