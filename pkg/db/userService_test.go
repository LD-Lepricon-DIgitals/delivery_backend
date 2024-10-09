package db_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	dbServ "github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	userService := dbServ.NewUserService(sqlxDB)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO users").
		WithArgs("login_test", "hashed_password_test").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec("INSERT INTO users_info").
		WithArgs(1, "123456789", "John", "Doe", "123 Test St").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	userId, err := userService.CreateUser("login_test", "John", "Doe", "123 Test St", "123456789", "hashed_password_test")

	assert.NoError(t, err)
	assert.Equal(t, 1, userId)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}
