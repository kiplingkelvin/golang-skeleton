package postgres

import (
	"fmt"
	"kiplingkelvin/golang-skeleton/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
)

type Config struct {
	DatabaseName string `envconfig:"DB_DATABASE_NAME" required:"true"`
	User         string `envconfig:"DB_DATABASE_USER" required:"true"`
	Password     string `envconfig:"DB_DATABASE_PASSWORD" required:"true"`
	Host         string `envconfig:"DB_DATABASE_HOST" required:"true"`
	Port         uint32 `envconfig:"DB_DATABASE_PORT" required:"false" default:"5432" `
}

type Postgres struct {
	db     *gorm.DB
	config *Config
}

func NewPostgres(config *Config) *Postgres {
	return &Postgres{
		config: config,
	}
}

func (dao *Postgres) Connect() (db *gorm.DB, err error) {
    logrus.Info("Connecting to postgres DB")
	
	db, err = gorm.Open(postgres.Open(dao.getConnectionString()), nil)
	if err != nil {
		return nil, err
	}
	
	dao.db = db
	return db, nil
}

func (dao *Postgres) getConnectionString() string {
	fmt.Println(dao.config)
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

		//Migrations
		logrus.Info("Running migrations")
		db.AutoMigrate(&models.HealthCheck{})
		dao.db = db
	}
	return dao.db, nil
}

