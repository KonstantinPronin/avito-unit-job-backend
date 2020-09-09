package repository

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/model"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Transaction struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (t *Transaction) Add(tran *model.Transaction) (*model.Transaction, error) {
	panic("implement me")
}

func (t *Transaction) Get(tid uint64) (*model.Transaction, error) {
	panic("implement me")
}

func (t *Transaction) GetOrderByDate(uid uint64, offset, limit uint) {
	panic("implement me")
}

func (t *Transaction) GetOrderBySum(uid uint64, offset, limit uint) {
	panic("implement me")
}

func NewTransaction(db *gorm.DB, logger *zap.Logger) transaction.Repository {
	return &Transaction{
		db:     db,
		logger: logger,
	}
}
