package usecase

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/currency"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/currency/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"strings"
	"testing"
)

func setUp(t *testing.T) (*mock.MockRepository, currency.Usecase) {
	ctrl := gomock.NewController(t)
	rep := mock.NewMockRepository(ctrl)
	usecase := &Currency{rep: rep, logger: zap.NewExample()}

	return rep, usecase
}

func TestCurrency_Exchange(t *testing.T) {
	cur := "USD"
	rate := float32(0.5)
	val := 70.0
	expected := float64(rate) * val
	rep, usecase := setUp(t)

	rep.EXPECT().GetByCurrency(cur).Return(rate, nil)

	actual, err := usecase.Exchange(val, strings.ToLower(cur))

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
