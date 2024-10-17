package db_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/models"
	dbServ "github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDishService_GetDishes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dishService := dbServ.NewDishService(sqlxDB)

	type args struct{}

	type mockBehavior func()
	testTable := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		want         []models.Dish
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{},
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "dish_name", "dish_description", "dish_price", "dish_weight", "dish_photo", "dish_rating", "category_name"}).
					AddRow(1, "Plov", "desc", 100.0, 100.1, "url", 2.2, "category")
				mock.ExpectQuery("SELECT dishes.id, dishes.dish_name, dishes.dish_description, dishes.dish_price, dishes.dish_weight, dishes.dish_photo, dishes.dish_rating, dish_categories.category_name FROM dishes LEFT JOIN dish_categories ON dishes.dish_category=dish_categories.id;").
					WillReturnRows(rows).
					RowsWillBeClosed()
			},
			want: []models.Dish{
				{
					Id:          1,
					Name:        "Plov",
					Description: "desc",
					Price:       100.0,
					Weight:      100.1,
					PhotoUrl:    "url",
					Rating:      2.2,
					Category:    "category"},
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()
			dishes, err := dishService.GetDishes()
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, testCase.want, dishes)
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
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
		dishCategory    int
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
				dishCategory:    1,
			},
			id: 1,
			mockBehavior: func(args args) {
				mock.ExpectQuery("INSERT INTO dishes").
					WithArgs(args.dishName, args.dishPrice, args.dishWeight, args.dishDescription, args.dishPhoto, args.dishCategory).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)
			got, err := dishService.AddDish(testCase.args.dishName, testCase.args.dishPrice, testCase.args.dishWeight, testCase.args.dishDescription, testCase.args.dishPhoto, testCase.args.dishCategory)
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

	type args struct {
		id int
	}
	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		wantError    bool
	}{
		{
			name: "OK",
			args: args{
				id: 1,
			},
			mockBehavior: func(args args) {
				mock.ExpectExec("DELETE FROM dishes WHERE id=\\$1;").
					WithArgs(args.id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantError: false,
		},
		{
			name: "Args without id",
			args: args{
				id: 1,
			},
			mockBehavior: func(args args) {
				mock.ExpectExec("DELETE FROM dishes WHERE id=\\$1;").
					WithArgs().
					WillReturnError(errors.New("no Id in args"))
			},
			wantError: true,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)
			err := dishService.DeleteDish(testCase.args.id)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				err = mock.ExpectationsWereMet()
				assert.NoError(t, err)
			}
		})
	}
}

func TestDishService_ChangeDish(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dishService := dbServ.NewDishService(sqlxDB)

	type args struct {
		id              int
		dishName        string
		dishPrice       float64
		dishWeight      float64
		dishDescription string
		dishPhoto       string
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				id:              1,
				dishName:        "Some name",
				dishPrice:       101.0,
				dishWeight:      10.0,
				dishDescription: "Some description",
				dishPhoto:       "photo url",
			},
			mockBehavior: func(args args) {
				mock.ExpectQuery("SELECT COUNT\\(1\\) FROM dishes WHERE id=\\$1;").
					WithArgs(args.id).
					WillReturnRows(sqlmock.NewRows([]string{"res"}).AddRow(1))
				mock.ExpectExec("UPDATE dishes SET dish_name=\\$1, dish_price=\\$2, dish_weight=\\$3, dish_description=\\$4, dish_photo=\\$5 WHERE id=\\$6;").
					WithArgs(args.dishName, args.dishPrice, args.dishWeight, args.dishDescription, args.dishPhoto, args.id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "Error after checking",
			args: args{
				id:              1,
				dishName:        "Some name",
				dishPrice:       101.0,
				dishWeight:      10.0,
				dishDescription: "Some description",
				dishPhoto:       "photo url",
			},
			mockBehavior: func(args args) {
				mock.ExpectQuery("SELECT COUNT\\(1\\) FROM dishes WHERE id=\\$1;").
					WithArgs(args.id).
					WillReturnRows(sqlmock.NewRows([]string{"res"}).AddRow(0)).WillReturnError(errors.New("error checking dish count"))
			},
			wantErr: true,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)
			err := dishService.ChangeDish(testCase.args.id, testCase.args.dishName, testCase.args.dishPrice, testCase.args.dishWeight, testCase.args.dishDescription, testCase.args.dishPhoto)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				err = mock.ExpectationsWereMet()
				assert.NoError(t, err)
			}

		})
	}
}
