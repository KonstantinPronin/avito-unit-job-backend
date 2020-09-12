package usecase

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/model"
	errs "github.com/KonstantinPronin/avito-unit-job-backend/pkg/model"
	"go.uber.org/zap"
	"time"
)

type Transaction struct {
	pageSize uint
	rep      transaction.Repository
	logger   *zap.Logger
}

func (t *Transaction) Add(tr *model.Transaction) (*model.Transaction, error) {
	if tr.From == 0 || tr.To == 0 || tr.Sum == 0 {
		return nil, errs.NewInvalidArgument("wrong transaction parameters")
	}

	tr.Created = time.Now()

	return t.rep.Add(tr)
}

func (t *Transaction) Get(tid uint64) (*model.Transaction, error) {
	if tid == 0 {
		return nil, errs.NewInvalidArgument("wrong transaction id")
	}

	return t.rep.Get(tid)
}

func (t *Transaction) GetOrderByDate(uid uint64, page uint, desc bool) (model.History, error) {
	if uid == 0 {
		return nil, errs.NewInvalidArgument("wrong transaction id")
	}

	offset := page * t.pageSize
	limit := offset + t.pageSize

	return t.rep.GetOrderByDate(uid, offset, limit, desc)
}

func (t *Transaction) GetOrderBySum(uid uint64, page uint, desc bool) (model.History, error) {
	if uid == 0 {
		return nil, errs.NewInvalidArgument("wrong transaction id")
	}

	offset := page * t.pageSize
	limit := offset + t.pageSize

	return t.rep.GetOrderBySum(uid, offset, limit, desc)
}

func NewTransaction(pageSize uint, rep transaction.Repository, logger *zap.Logger) transaction.Usecase {
	return &Transaction{
		pageSize: pageSize,
		rep:      rep,
		logger:   logger,
	}
}
