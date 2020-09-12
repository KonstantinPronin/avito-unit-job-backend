package http

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/model"
	"github.com/KonstantinPronin/avito-unit-job-backend/pkg/middleware"
	response "github.com/KonstantinPronin/avito-unit-job-backend/pkg/model"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type TransactionHandler struct {
	usecase   transaction.Usecase
	server    *echo.Echo
	logger    *zap.Logger
	sanitizer *bluemonday.Policy
}

func NewTransactionHandler(
	usecase transaction.Usecase,
	server *echo.Echo,
	logger *zap.Logger,
	sanitizer *bluemonday.Policy) {
	handler := &TransactionHandler{
		usecase:   usecase,
		server:    server,
		logger:    logger,
		sanitizer: sanitizer,
	}

	server.POST("/transaction", handler.AddTransaction, middleware.ParseErrors)
	server.GET("/transaction/:tid", handler.GetTransaction, middleware.ParseErrors)
	server.GET("/history/:uid", handler.GetHistory, middleware.ParseErrors)
}

func (th *TransactionHandler) AddTransaction(ctx echo.Context) error {
	tr, err := th.unmarshalFromReader(ctx.Request().Body)
	if err != nil {
		return err
	}

	tr, err = th.usecase.Add(tr)
	if err != nil {
		return err
	}

	return response.NewResponse(http.StatusOK, tr).WriteToCtx(ctx)
}

func (th *TransactionHandler) GetTransaction(ctx echo.Context) error {
	tid, err := middleware.GetIdParam(ctx, "tid")
	if err != nil {
		return err
	}

	tr, err := th.usecase.Get(tid)
	if err != nil {
		return err
	}

	return response.NewResponse(http.StatusOK, tr).WriteToCtx(ctx)
}

func (th *TransactionHandler) GetHistory(ctx echo.Context) error {
	uid, err := middleware.GetIdParam(ctx, "uid")
	if err != nil {
		return err
	}

	var history model.History
	params := NewQueryParams().init(ctx.QueryParams())

	switch params.Order {
	case Sum:
		history, err = th.usecase.GetOrderBySum(uid, params.Page, params.Desc)
	default:
		history, err = th.usecase.GetOrderByDate(uid, params.Page, params.Desc)
	}

	if err != nil {
		return err
	}

	return response.NewResponse(http.StatusOK, history).WriteToCtx(ctx)
}

func (th *TransactionHandler) sanitize(tr *model.Transaction) {
	tr.Comment = th.sanitizer.Sanitize(tr.Comment)
}

func (th *TransactionHandler) unmarshalFromReader(r io.Reader) (*model.Transaction, error) {
	tr := new(model.Transaction)

	if err := easyjson.UnmarshalFromReader(r, tr); err != nil {
		return nil, response.NewInvalidArgument("wrong request body format")
	}
	th.sanitize(tr)

	return tr, nil
}
