package repository

import (
	"bytes"
	"fmt"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/currency"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/currency/model"
	"github.com/KonstantinPronin/avito-unit-job-backend/pkg/constants"
	"github.com/KonstantinPronin/avito-unit-job-backend/pkg/http"
	errs "github.com/KonstantinPronin/avito-unit-job-backend/pkg/model"
	"github.com/alicebob/miniredis/v2"
	"github.com/elliotchance/redismock/v7"
	"github.com/go-redis/redis/v7"
	"github.com/golang/mock/gomock"
	"github.com/mailru/easyjson"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"io/ioutil"
	outhttp "net/http"
	"strconv"
	"testing"
	"time"
)

func setUp(t *testing.T) (*http.MockGetter, *redismock.ClientMock, currency.Repository) {
	ctrl := gomock.NewController(t)
	getter := http.NewMockGetter(ctrl)

	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub redis connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mock := redismock.NewNiceMock(client)

	return getter, mock,
		&Currency{base: "test", conn: mock, logger: zap.NewExample(), client: &outhttp.Client{}}
}

func TestCurrency_GetByCurrency_FromCache(t *testing.T) {
	cur := "USD"
	expected := 70.0
	_, rd, rep := setUp(t)

	rd.On("Get", cur).
		Return(redis.NewStringResult(strconv.FormatFloat(expected, 'f', -1, 32), nil))

	actual, err := rep.GetByCurrency(cur)

	assert.Nil(t, err)
	assert.Equal(t, float32(expected), actual)
}

func TestCurrency_GetByCurrency_NotFound(t *testing.T) {
	cur := "USD"
	_, _, rep := setUp(t)

	_, err := rep.GetByCurrency(cur)

	assert.NotNil(t, err)
	assert.Equal(t, err, errs.NewNotFoundError(fmt.Sprintf("Cannot find currency = %s", cur)))
}

func TestCurrency_UpdateCache(t *testing.T) {
	cur := "USD"
	expected := float32(70.0)

	gtr, rd, _ := setUp(t)
	url := fmt.Sprintf("%s?base=%s", constants.ExchangeRateURL, "test")
	rep := &Currency{base: "test", conn: rd, logger: zap.NewExample(), client: gtr}

	rates := model.ExchangeRates{Rates: map[string]float32{cur: expected}}
	body, err := easyjson.Marshal(rates)
	if err != nil {
		t.Errorf("unexpected error: '%s'", err)
	}

	resp := &outhttp.Response{StatusCode: outhttp.StatusOK, Body: ioutil.NopCloser(bytes.NewReader(body))}

	gtr.EXPECT().Get(url).Return(resp, nil)
	rd.On("Set", cur, rates.Rates[cur], 24*time.Hour).Return(redis.NewStatusResult("", nil))

	actual, err := rep.GetByCurrency(cur)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
