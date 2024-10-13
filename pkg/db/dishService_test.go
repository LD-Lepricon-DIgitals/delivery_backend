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

	type args struct {
		dishName        string
		dishPrice       float64
		dishWeight      float64
		dishDescription string
		dishPhoto       string
	}
	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		id           int
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				dishName:        "test",
				dishPrice:       10.1,
				dishWeight:      100.0,
				dishDescription: "test",
				dishPhoto:       "url",
			},
			id: 1,
			mockBehavior: func(args args) {
				mock.ExpectQuery("INSERT INTO dishes").
					WithArgs(args.dishName, args.dishPrice, args.dishWeight, args.dishDescription, args.dishPhoto).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)
			got, err := dishService.AddDish(testCase.args.dishName, testCase.args.dishPrice, testCase.args.dishWeight, testCase.args.dishDescription, testCase.args.dishPhoto)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.id, got)
				err = mock.ExpectationsWereMet()
				assert.NoError(t, err)
			}
		})
	}
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
