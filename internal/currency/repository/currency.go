package repository

import (
	"fmt"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/currency"
	"github.com/KonstantinPronin/avito-unit-job-backend/internal/currency/model"
	"github.com/KonstantinPronin/avito-unit-job-backend/pkg/constants"
	"github.com/KonstantinPronin/avito-unit-job-backend/pkg/http"
	errs "github.com/KonstantinPronin/avito-unit-job-backend/pkg/model"
	"github.com/go-redis/redis/v7"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	outhttp "net/http"
	"strconv"
	"time"
)

type Currency struct {
	base   string
	conn   redis.Cmdable
	client http.Getter
	logger *zap.Logger
}

func (c *Currency) GetByCurrency(cur string) (float32, error) {
	cache, err := c.conn.Get(cur).Result()

	if err == nil {
		rate, err := strconv.ParseFloat(cache, 32)

		if err != nil {
			return 0, err
		}

		return float32(rate), nil
	}

	if err == redis.Nil {
		rates, err := c.UpdateCache()
		if err != nil {
			return 0, err
		}

		rate, ok := rates[cur]
		if !ok {
			return 0, errs.NewNotFoundError(fmt.Sprintf("Cannot find currency = %s", cur))
		}

		return rate, nil
	}

	return 0, err
}

func (c *Currency) UpdateCache() (map[string]float32, error) {
	rates, err := c.getExchangeRates()

	if err != nil {
		c.logger.Warn(fmt.Sprintf("error while requesting new exchange rates: %s", err.Error()))
		return nil, err
	}

	for cur, rate := range rates {
		err = c.conn.Set(cur, rate, 24*time.Hour).Err()
		if err != nil {
			c.logger.Warn(fmt.Sprintf("error while saving exchange rates: %s", err.Error()))
			return nil, err
		}
	}

	return rates, nil
}

func (c *Currency) getExchangeRates() (map[string]float32, error) {
	if c.base == "" {
		c.base = constants.DefaultBaseCurrency
	}

	url := fmt.Sprintf("%s?base=%s", constants.ExchangeRateURL, c.base)

	resp, err := c.client.Get(url)
	if resp != nil {
		defer func() {
			if err = resp.Body.Close(); err != nil {
				c.logger.Error(fmt.Sprintf("Resource leakes %s", err.Error()))
			}
		}()
	}

	if err != nil {
		return nil, err
	}

	exchange := new(model.ExchangeRates)
	if err = easyjson.UnmarshalFromReader(resp.Body, exchange); err != nil {
		return nil, err
	}

	return exchange.Rates, nil
}

func NewCurrency(base string,
	conn redis.Cmdable,
	logger *zap.Logger) currency.Repository {
	return &Currency{
		base:   base,
		conn:   conn,
		logger: logger,
		client: &outhttp.Client{},
	}
}
