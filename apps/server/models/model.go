package models

import "time"

type LoginData struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type Tasks struct {
	Id   uint   `json:"id"`
	Date string `json:"date"`
	Text string `json:"text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
