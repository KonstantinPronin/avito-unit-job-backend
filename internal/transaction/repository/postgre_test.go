package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/model"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

var (
	test = model.Transaction{
		ID:      1,
		From:    1,
		To:      1,
		Created: time.Time{},
		Sum:     10,
		Comment: "test",
	}

	uid    = uint64(1)
	offset = uint(1)
	limit  = uint(1)
)

func beforeTest(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, transaction.Repository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	conn, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
	}

	rep := NewTransaction(conn, zap.NewExample())
	return conn, mock, rep
}

func TestTransaction_Add(t *testing.T) {
	db, mock, rep := beforeTest(t)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO (.*)transactions`).
		WithArgs(test.ID, test.From, test.To, test.Created, test.Sum, test.Comment).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(test.ID))
	mock.ExpectCommit()

	tr, err := rep.Add(&test)

	assert.Nil(t, err)
	assert.Equal(t, &test, tr)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTransaction_Get(t *testing.T) {
	db, mock, rep := beforeTest(t)
	defer db.Close()

	mock.ExpectQuery(`SELECT (\*) FROM (.*)transactions(.*) WHERE (.*)id(.*) LIMIT 1`).
		WithArgs(test.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "from_user", "to_user", "created", "sum", "comment"}).
			AddRow(test.ID, test.From, test.To, test.Created, test.Sum, test.Comment))

	tr, err := rep.Get(test.ID)

	assert.Nil(t, err)
	assert.Equal(t, &test, tr)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTransaction_GetOrderByDateAsc(t *testing.T) {
	db, mock, rep := beforeTest(t)
	defer db.Close()

	mock.ExpectQuery(`SELECT (\*) FROM (.*)transactions(.*) 
				WHERE (.*)from_user(.*) OR (.*)to_user(.*) ORDER BY created asc LIMIT (.*) OFFSET (.*)`).
		WithArgs(uid, uid).
		WillReturnRows(sqlmock.NewRows([]string{"id", "from_user", "to_user", "created", "sum", "comment"}).
			AddRow(test.ID, test.From, test.To, test.Created, test.Sum, test.Comment))

	history, err := rep.GetOrderByDate(test.ID, offset, limit, false)

	assert.Nil(t, err)
	assert.Equal(t, test, history[0])

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTransaction_GetOrderByDateDesc(t *testing.T) {
	db, mock, rep := beforeTest(t)
	defer db.Close()

	mock.ExpectQuery(`SELECT (\*) FROM (.*)transactions(.*) 
				WHERE (.*)from_user(.*) OR (.*)to_user(.*) ORDER BY created desc LIMIT (.*) OFFSET (.*)`).
		WithArgs(uid, uid).
		WillReturnRows(sqlmock.NewRows([]string{"id", "from_user", "to_user", "created", "sum", "comment"}).
			AddRow(test.ID, test.From, test.To, test.Created, test.Sum, test.Comment))

	history, err := rep.GetOrderByDate(test.ID, offset, limit, true)

	assert.Nil(t, err)
	assert.Equal(t, test, history[0])

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTransaction_GetOrderBySumAsc(t *testing.T) {
	db, mock, rep := beforeTest(t)
	defer db.Close()

	mock.ExpectQuery(`SELECT (\*) FROM (.*)transactions(.*) 
				WHERE (.*)from_user(.*) OR (.*)to_user(.*) ORDER BY sum asc LIMIT (.*) OFFSET (.*)`).
		WithArgs(uid, uid).
		WillReturnRows(sqlmock.NewRows([]string{"id", "from_user", "to_user", "created", "sum", "comment"}).
			AddRow(test.ID, test.From, test.To, test.Created, test.Sum, test.Comment))

	history, err := rep.GetOrderBySum(test.ID, offset, limit, false)

	assert.Nil(t, err)
	assert.Equal(t, test, history[0])

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTransaction_GetOrderBySumDesc(t *testing.T) {
	db, mock, rep := beforeTest(t)
	defer db.Close()

	mock.ExpectQuery(`SELECT (\*) FROM (.*)transactions(.*) 
				WHERE (.*)from_user(.*) OR (.*)to_user(.*) ORDER BY sum desc LIMIT (.*) OFFSET (.*)`).
		WithArgs(uid, uid).
		WillReturnRows(sqlmock.NewRows([]string{"id", "from_user", "to_user", "created", "sum", "comment"}).
			AddRow(test.ID, test.From, test.To, test.Created, test.Sum, test.Comment))

	history, err := rep.GetOrderBySum(test.ID, offset, limit, true)

	assert.Nil(t, err)
	assert.Equal(t, test, history[0])

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
