package tests

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := SetupMockDB()
	assert.NoError(t, err, "Failed to setup mock database")

	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Expectation for creating a user
	mock.ExpectBegin()
	expectedSql := `INSERT INTO "users" \("name","email","password","created_at","updated_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5\) RETURNING "id"`
	mock.ExpectQuery(expectedSql).
		WithArgs(user.Name, user.Email, "", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	err = repo.CreateUser(user)
	assert.NoError(t, err, "Failed to create user")
	assert.NoError(t, mock.ExpectationsWereMet(), "Mock expectations were not met")
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := SetupMockDB()
	assert.NoError(t, err, "Failed to setup mock database")

	repo := repositories.NewUserRepository(db)

	email := "john@example.com"
	user := &models.User{
		ID:       1,
		Name:     "John Doe",
		Email:    email,
		Password: "",
	}

	// Expectation for retrieving a user by email
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.Password, time.Now(), time.Now())
	expectedSql := `SELECT \* FROM "users" WHERE email = \$1 ORDER BY "users"\."id" LIMIT \$2`
	mock.ExpectQuery(expectedSql).WithArgs(email, 1).WillReturnRows(rows)

	result, err := repo.GetUserByEmail(user.Email)
	assert.NoError(t, err, "Failed to retrieve user by email")
	assert.Equal(t, user.Name, result.Name, "User name does not match")
	assert.Equal(t, user.Email, result.Email, "User email does not match")
	assert.NoError(t, mock.ExpectationsWereMet(), "Mock expectations were not met")
}
