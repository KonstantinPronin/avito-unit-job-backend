package http

import (
	"bytes"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/mock"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/model"
	mock2 "github.com/KonstantinPronin/avito-unit-job-backend/pkg/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"strings"
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

	wrongID   = "test"
	correctID = "1"
)

func setUp(t *testing.T) (
	*mock.MockUsecase,
	*mock2.MockContext,
	*TransactionHandler) {
	ctrl := gomock.NewController(t)

	usecase := mock.NewMockUsecase(ctrl)
	handler := &TransactionHandler{usecase: usecase, logger: zap.NewExample(), sanitizer: bluemonday.UGCPolicy()}

	ctx := mock2.NewMockContext(ctrl)
	writer := mock2.NewMockResponseWriter(ctrl)

	response := echo.NewResponse(writer, echo.New())
	ctx.EXPECT().Response().Return(response).AnyTimes()
	writer.EXPECT().Write(gomock.Any())

	return usecase, ctx, handler
}

func TestTransactionHandler_AddTransaction_EmptyBody(t *testing.T) {
	_, ctx, handler := setUp(t)

	request, err := http.NewRequest("", "", strings.NewReader(""))
	if err != nil {
		t.Errorf("unexpected error: '%s'", err)
	}

	ctx.EXPECT().Request().Return(request).AnyTimes()

	err = handler.AddTransaction(ctx)

	assert.NotNil(t, err)
}

func TestTransactionHandler_AddTransaction(t *testing.T) {
	usecase, ctx, handler := setUp(t)

	body, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf("unexpected error: '%s'", err)
	}

	request, err := http.NewRequest("", "", bytes.NewReader(body))
	if err != nil {
		t.Errorf("unexpected error: '%s'", err)
	}

	ctx.EXPECT().Request().Return(request).AnyTimes()
	usecase.EXPECT().Add(&test).Return(&test, nil)

	err = handler.AddTransaction(ctx)

	assert.Nil(t, err)
}

func TestTransactionHandler_GetTransaction_WrongId(t *testing.T) {
	_, ctx, handler := setUp(t)

	ctx.EXPECT().Param(gomock.Any()).Return(wrongID)

	err := handler.GetTransaction(ctx)

	assert.NotNil(t, err)
}

func TestTransactionHandler_GetTransaction(t *testing.T) {
	usecase, ctx, handler := setUp(t)

	ctx.EXPECT().Param(gomock.Any()).Return(correctID)
	usecase.EXPECT().Get(test.ID).Return(&test, nil)

	err := handler.GetTransaction(ctx)

	assert.Nil(t, err)
}

func TestTransactionHandler_GetHistory_WrongID(t *testing.T) {
	_, ctx, handler := setUp(t)

	ctx.EXPECT().Param(gomock.Any()).Return(wrongID)

	err := handler.GetHistory(ctx)

	assert.NotNil(t, err)
}

func TestTransactionHandler_GetHistory_DateOrder(t *testing.T) {
	usecase, ctx, handler := setUp(t)

	ctx.EXPECT().Param(gomock.Any()).Return(correctID)
	ctx.EXPECT().QueryParams().Return(map[string][]string{"order": {"date"}})
	usecase.EXPECT().GetOrderByDate(test.ID, uint(0), false).Return([]model.Transaction{test}, nil)

	err := handler.GetHistory(ctx)

	assert.Nil(t, err)
}

func TestTransactionHandler_GetHistory_SumOrder(t *testing.T) {
	usecase, ctx, handler := setUp(t)

	ctx.EXPECT().Param(gomock.Any()).Return(correctID)
	ctx.EXPECT().QueryParams().Return(map[string][]string{"order": {"sum"}})
	usecase.EXPECT().GetOrderBySum(test.ID, uint(0), false).Return([]model.Transaction{test}, nil)

	err := handler.GetHistory(ctx)

	assert.Nil(t, err)
}
