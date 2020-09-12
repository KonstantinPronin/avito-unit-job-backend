package repository

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/transaction/model"
	errs "github.com/KonstantinPronin/avito-unit-job-backend/pkg/model"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Transaction struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (t *Transaction) Add(tran *model.Transaction) (*model.Transaction, error) {
	return tran, t.db.Table("job.transactions").Create(tran).Error
}

func (t *Transaction) Get(tid uint64) (*model.Transaction, error) {
	tr := new(model.Transaction)

	if err := t.db.Table("job.transactions").Where("id = ?", tid).First(tr).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errs.NewNotFoundError(err.Error())
		}

		return nil, err
	}

	return tr, nil
}

func (t *Transaction) GetOrderByDate(uid uint64, offset, limit uint, desc bool) (model.History, error) {
	order := "created asc"
	if desc {
		order = "created desc"
	}

	return t.GetOrderBy(uid, offset, limit, order)
}

func (t *Transaction) GetOrderBySum(uid uint64, offset, limit uint, desc bool) (model.History, error) {
	order := "sum asc"
	if desc {
		order = "sum desc"
	}

	return t.GetOrderBy(uid, offset, limit, order)
}

func (t *Transaction) GetOrderBy(uid uint64, offset, limit uint, order string) (model.History, error) {
	var result model.History

	err := t.db.Table("job.transactions").
		Where("from_user = ?", uid).Or("to_user = ?", uid).
		Offset(offset).Limit(limit).Order(order).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewTransaction(db *gorm.DB, logger *zap.Logger) transaction.Repository {
	return &Transaction{
		db:     db,
		logger: logger,
	}
}
