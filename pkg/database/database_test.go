package database_test

import (
	"testing"

	"go-rest-api/config"
	"go-rest-api/pkg/database"
	"gorm.io/gorm/logger"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setupConfig sets up a test configuration
func setupConfig() {
	v := viper.New()
	v.Set("database.host", "localhost")
	v.Set("database.port", 5432)
	v.Set("database.user", "test_user")
	v.Set("database.dbname", "test_db")
	v.Set("database.password", "test_password")
	config.Config = v
}

func TestInitDB(t *testing.T) {
	// Set up test configuration
	setupConfig()

	// Initialize sqlmock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	// Set expectation for the mock query
	mock.ExpectQuery("^SELECT 1$").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	// Set expectation for closing the database
	mock.ExpectClose()

	database.SetGormOpen(func(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error) {
		return gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	})
	defer database.SetGormOpen(gorm.Open) // Reset after test

	dsn := database.GetDSN()

	// Inject sqlmock into InitDB for testing
	mockDB, err := database.InitDB(dsn)
	assert.NoError(t, err)
	assert.NotNil(t, mockDB)

	// Trigger an interaction to fulfill the expectation
	var result int
	err = mockDB.Raw("SELECT 1").Scan(&result).Error
	assert.NoError(t, err)

	// Close the mock database connection
	sqlDB, err := mockDB.DB()
	assert.NoError(t, err)
	err = sqlDB.Close()
	assert.NoError(t, err)

	// Ensure all expectations are met (before closing db)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "There were unmet expectations")
}

func TestInitDB_FailedConnection(t *testing.T) {
	setupConfig()

	// Mock a bad connection scenario
	database.SetGormOpen(func(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error) {
		return nil, gorm.ErrInvalidDB
	})
	defer database.SetGormOpen(gorm.Open) // Reset after test

	_, err := database.InitDB(database.GetDSN())
	assert.Error(t, err)
}
