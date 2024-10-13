package db_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	dbServ "github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDishService_AddDish(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dishService := dbServ.NewDishService(sqlxDB)

	mock.ExpectQuery("INSERT INTO dishes").
		WithArgs("dish_name_test", 100.0, 100.1, "dish_description_test", "dish_photo_test").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := dishService.AddDish("dish_name_test", 100.0, 100.1, "dish_description_test", "dish_photo_test")
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
func TestDishService_DeleteDish(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dishService := dbServ.NewDishService(sqlxDB)

	mock.ExpectExec("DELETE FROM dishes WHERE id=\\$1;").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = dishService.DeleteDish(1)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
