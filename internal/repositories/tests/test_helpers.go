package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setupMockDB initializes a mock database and returns the gorm.DB instance and sqlmock.Sqlmock for testing.
func SetupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        "sqlmock_db_0",
		Conn:       mockDB,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}
