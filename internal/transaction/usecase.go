package transaction

import "github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/model"

type Usecase interface {
	Add(tr *model.Transaction) (*model.Transaction, error)
	Get(tid uint64) (*model.Transaction, error)
	GetOrderByDate(uid uint64, page uint, desc bool) (model.History, error)
	GetOrderBySum(uid uint64, page uint, desc bool) (model.History, error)
}
