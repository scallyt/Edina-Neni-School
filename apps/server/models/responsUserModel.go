package models

type responsUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
