package http

import (
	"net/url"
	"strconv"
	"strings"
)

type Order int

const (
	Date Order = iota
	Sum
)

type QueryParams struct {
	Order Order
	Page  uint
	Desc  bool
}

func NewQueryParams() *QueryParams {
	return &QueryParams{
		Order: Date,
		Page:  0,
		Desc:  false,
	}
}

func (qp *QueryParams) Init(params url.Values) *QueryParams {
	order := params.Get("order")
	if strings.EqualFold(order, "sum") {
		qp.Order = Sum
	}

	page, err := strconv.Atoi(params.Get("page"))
	if err == nil && page > 0 {
		qp.Page = uint(page)
	}

	desc := params.Get("desc")
	if strings.EqualFold(desc, "true") {
		qp.Desc = true
	}

	return qp
}
