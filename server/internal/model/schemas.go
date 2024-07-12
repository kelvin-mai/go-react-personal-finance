package model

import "time"

type User struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
}

type Category struct {
	Id        string     `json:"id"`
	UserId    string     `db:"user_id" json:"user_id"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

type Transaction struct {
	Id         string     `json:"id"`
	UserId     string     `db:"user_id" json:"userId"`
	CategoryId string     `db:"category_id" json:"categoryId"`
	Title      string     `json:"title"`
	Amount     float32    `json:"amount"`
	Currency   string     `json:"currency"`
	Type       string     `json:"type"`
	CreatedAt  time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updatedAt"`
}
