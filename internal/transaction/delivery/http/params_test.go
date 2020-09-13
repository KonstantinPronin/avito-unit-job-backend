package http

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var testCase = map[QueryParams]map[string][]string{
	QueryParams{
		Order: Date,
		Page:  0,
		Desc:  false,
	}: {},
	QueryParams{
		Order: Sum,
		Page:  0,
		Desc:  false,
	}: {
		"order": {"sum"},
	},
	QueryParams{
		Order: Sum,
		Page:  10,
		Desc:  false,
	}: {
		"order": {"sum"},
		"page":  {"10"},
	},
	QueryParams{
		Order: Sum,
		Page:  10,
		Desc:  true,
	}: {
		"order": {"sum"},
		"page":  {"10"},
		"desc":  {"true"},
	},
}

func TestQueryParams_Init(t *testing.T) {
	for expected, params := range testCase {
		actual := NewQueryParams().Init(params)
		assert.Equal(t, *actual, expected)
	}
}
