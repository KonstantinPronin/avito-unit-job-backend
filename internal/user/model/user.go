package model

type User struct {
	ID       uint64  `json:"id" gorm:"primary_key, column:id"`
	Username string  `json:"username" gorm:"column:username"`
	Balance  float64 `json:"balance" gotm:"column:balance"`
}
