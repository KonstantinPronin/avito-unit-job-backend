package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/user/model"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

var (
	test = model.User{
		ID:       1,
		Username: "test",
		Balance:  10,
	}
)

func beforeTest(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	conn, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
	}

	return conn, mock
}

func TestUser_Add(t *testing.T) {
	db, mock := beforeTest(t)
	defer db.Close()

	r := NewUser(db, zap.NewExample())

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO (.*)users`).
		WithArgs(test.ID, test.Username, test.Balance).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(test.ID))
	mock.ExpectCommit()

	usr, err := r.Add(&test)

	assert.Nil(t, err)
	assert.Equal(t, &test, usr)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_Get_NotFound(t *testing.T) {
	db, mock := beforeTest(t)
	defer db.Close()

	r := NewUser(db, zap.NewExample())

	mock.ExpectQuery(`SELECT (\*) FROM (.*)users(.*) WHERE (.*)id(.*) LIMIT 1`).
		WithArgs(test.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "balance"}))

	_, err := r.Get(test.ID)

	assert.NotNil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_Get(t *testing.T) {
	db, mock := beforeTest(t)
	defer db.Close()

	r := NewUser(db, zap.NewExample())

	mock.ExpectQuery(`SELECT (\*) FROM (.*)users(.*) WHERE (.*)id(.*) LIMIT 1`).
		WithArgs(test.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "balance"}).
		AddRow(test.ID, test.Username, test.Balance))

	usr, err := r.Get(test.ID)

	assert.Nil(t, err)
	assert.Equal(t, &test, usr)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_Update(t *testing.T) {
	db, mock := beforeTest(t)
	defer db.Close()

	r := NewUser(db, zap.NewExample())

	mock.ExpectQuery(`SELECT (\*) FROM (.*)users(.*) WHERE (.*)id(.*) LIMIT 1`).
		WithArgs(test.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "balance"}).
		AddRow(test.ID, test.Username, test.Balance))
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE (.*)users(.*) SET (.*)username(.*) WHERE (.*)id(.*)").
		WithArgs(test.Username, test.Balance, test.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	usr, err := r.Update(test.ID, &test)

	assert.Nil(t, err)
	assert.Equal(t, &test, usr)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_Contains(t *testing.T) {
	db, mock := beforeTest(t)
	defer db.Close()

	r := NewUser(db, zap.NewExample())

	mock.ExpectQuery(`SELECT (\*) FROM (.*)users(.*) WHERE (.*)username(.*) LIMIT 1`).
		WithArgs(test.Username).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "balance"}).
		AddRow(test.ID, test.Username, test.Balance))

	ok := r.Contains(test.Username)

	assert.True(t, ok)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
