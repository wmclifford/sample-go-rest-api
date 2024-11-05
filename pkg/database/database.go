package database

import (
	"fmt"
	"log"

	"go-rest-api/config"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormOpen = gorm.Open

// GetDSN constructs the Data Source Name from the configuration.
func GetDSN() string {
	conf := config.Config
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable TimeZone=UTC",
		conf.GetString("database.host"),
		conf.GetInt("database.port"),
		conf.GetString("database.user"),
		conf.GetString("database.dbname"),
		conf.GetString("database.password"),
	)
}

// InitDB initializes and returns a database connection using the
// configuration provided in the config package.
func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gormOpen(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Printf("Could not connect to database: %v", err)
		return nil, err
	}
	return db, nil
}

// SetGormOpen is a setter for the gormOpen variable, used for dependency injection in tests.
func SetGormOpen(openFunc func(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error)) {
	gormOpen = openFunc
}
