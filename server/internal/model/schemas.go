package model

import "time"

type User struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
}
