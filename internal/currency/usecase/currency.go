package usecase

import (
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/currency"
	"go.uber.org/zap"
	"strings"
)

type Currency struct {
	rep    currency.Repository
	logger *zap.Logger
}

func (c *Currency) Exchange(value float64, cur string) (float64, error) {
	cur = strings.ToUpper(cur)

	rate, err := c.rep.GetByCurrency(cur)
	if err != nil {
		return 0, err
	}

	return float64(rate) * value, nil
}

func NewCurrency(rep currency.Repository,
	logger *zap.Logger) currency.Usecase {
	return &Currency{
		rep:    rep,
		logger: logger,
	}
}
