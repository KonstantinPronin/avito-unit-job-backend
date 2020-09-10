package model

import "time"

type Transaction struct {
	ID      uint64    `json:"-" gorm:"primary_key, column:id"`
	From    uint64    `json:"from" gorm:"column:from_user"`
	To      uint64    `json:"to" gorm:"column:to_user"`
	Created time.Time `json:"created" gorm:"column:created"`
	Sum     float64   `json:"sum" gorm:"column:sum"`
	Comment string    `json:"comment" gorm:"column:comment"`
}

//easyjson:json
type History []Transaction
