package transaction

import "github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/model"

type Repository interface {
	Add(tr *model.Transaction) (*model.Transaction, error)
	Get(tid uint64) (*model.Transaction, error)
	GetOrderByDate(uid uint64, offset, limit uint, desc bool) (model.History, error)
	GetOrderBySum(uid uint64, offset, limit uint, desc bool) (model.History, error)
}
