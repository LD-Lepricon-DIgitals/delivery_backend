package db_test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	dbServ "github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
	"github.com/jmoiron/sqlx"
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

	type args struct {
		login, name, surname, address, phoneNumber, password string
	}

	tests := []struct {
		name    string
		input   args
		mock    func(input args)
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				login:       "test login",
				name:        "test name",
				surname:     "test surname",
				address:     "test address",
				phoneNumber: "123456789",
				password:    "test password",
			},
			mock: func(input args) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO users").
					WithArgs(input.login, input.password).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectExec("INSERT INTO users_info").
					WithArgs(1, input.phoneNumber, input.name, input.surname, input.address).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Fail to insert into users",
			input: args{
				login:       "test login",
				name:        "test name",
				surname:     "test surname",
				address:     "test address",
				phoneNumber: "123456789",
				password:    "test password",
			},
			mock: func(input args) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO users").
					WithArgs(input.login, input.password).
					WillReturnError(fmt.Errorf("failed to insert into users"))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)
			_, err := userService.CreateUser(tt.input.login, tt.input.name, tt.input.surname, tt.input.address, tt.input.phoneNumber, tt.input.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %v", err)
			}
		})
	}
}

func TestUserService_ChangePassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	userService := dbServ.NewUserService(sqlxDB)

	tests := []struct {
		name    string
		userID  int
		newPass string
		mock    func(userID int, newPass string)
		wantErr bool
	}{
		{
			name:    "Ok",
			userID:  1,
			newPass: "new_password",
			mock: func(userID int, newPass string) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE users SET user_hashed_password = \\$1 WHERE id = \\$2").
					WithArgs(newPass, userID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name:    "User not found",
			userID:  2,
			newPass: "new_password",
			mock: func(userID int, newPass string) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE users SET user_hashed_password = \\$1 WHERE id = \\$2").
					WithArgs(newPass, userID).
					WillReturnResult(sqlmock.NewResult(0, 0))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:    "Error during update",
			userID:  3,
			newPass: "new_password",
			mock: func(userID int, newPass string) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE users SET user_hashed_password = \\$1 WHERE id = \\$2").
					WithArgs(newPass, userID).
					WillReturnError(fmt.Errorf("some error"))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:    "Error during commit",
			userID:  4,
			newPass: "new_password",
			mock: func(userID int, newPass string) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE users SET user_hashed_password = \\$1 WHERE id = \\$2").
					WithArgs(newPass, userID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit().WillReturnError(fmt.Errorf("commit error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.userID, tt.newPass)
			err := userService.ChangePassword(tt.userID, tt.newPass)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
			}
		})
	}
}
