package usecase

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/mock"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/model"
	"github.com/golang/mock/gomock"
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

	wrongID  = uint64(0)
	userID   = uint64(1)
	pageSize = uint(10)
	page     = uint(0)
	desc     = false
)

func setUp(t *testing.T) (*mock.MockRepository, transaction.Usecase) {
	ctrl := gomock.NewController(t)
	rep := mock.NewMockRepository(ctrl)
	usecase := NewTransaction(pageSize, rep, zap.NewExample())

	return rep, usecase
}

func TestTransaction_Add_WrongTransaction(t *testing.T) {
	_, usecase := setUp(t)
	var input []model.Transaction

	input = append(input, model.Transaction{
		From: 0,
		To:   1,
		Sum:  1,
	})
	input = append(input, model.Transaction{
		From: 1,
		To:   0,
		Sum:  1,
	})
	input = append(input, model.Transaction{
		From: 1,
		To:   1,
		Sum:  0,
	})

	for _, tr := range input {
		_, err := usecase.Add(&tr)
		assert.NotNil(t, err)
	}
}

func TestTransaction_Add(t *testing.T) {
	rep, usecase := setUp(t)

	rep.EXPECT().Add(&test).Return(&test, nil)

	tr, err := usecase.Add(&test)

	assert.Nil(t, err)
	assert.Equal(t, &test, tr)
}

func TestTransaction_Get_WrongId(t *testing.T) {
	_, usecase := setUp(t)

	_, err := usecase.Get(wrongID)

	assert.NotNil(t, err)
}

func TestTransaction_Get(t *testing.T) {
	rep, usecase := setUp(t)

	rep.EXPECT().Get(test.ID).Return(&test, nil)

	tr, err := usecase.Get(test.ID)

	assert.Nil(t, err)
	assert.Equal(t, &test, tr)
}

func TestTransaction_GetOrderByDate_WrongId(t *testing.T) {
	_, usecase := setUp(t)

	_, err := usecase.GetOrderByDate(wrongID, page, desc)

	assert.NotNil(t, err)
}

func TestTransaction_GetOrderByDate(t *testing.T) {
	rep, usecase := setUp(t)

	rep.EXPECT().GetOrderByDate(userID, page*pageSize, page*pageSize+pageSize, desc).Return([]model.Transaction{test}, nil)

	history, err := usecase.GetOrderByDate(userID, page, desc)

	assert.Nil(t, err)
	assert.Equal(t, test, history[0])
}

func TestTransaction_GetOrderBySum_WrongId(t *testing.T) {
	_, usecase := setUp(t)

	_, err := usecase.GetOrderBySum(wrongID, page, desc)

	assert.NotNil(t, err)
}

func TestTransaction_GetOrderBySum(t *testing.T) {
	rep, usecase := setUp(t)

	rep.EXPECT().GetOrderBySum(userID, page*pageSize, page*pageSize+pageSize, desc).Return([]model.Transaction{test}, nil)

	history, err := usecase.GetOrderBySum(userID, page, desc)

	assert.Nil(t, err)
	assert.Equal(t, test, history[0])
}
