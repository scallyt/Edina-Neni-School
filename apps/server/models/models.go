package models

import "github.com/scallyt/Edina-Neni-School/internal"

type LoginData struct {
	UserName string `json:"username"` // Capitalized field name
	Email    string `json:"email"`    // Capitalized field name
	Password string `json:"password"` // Capitalized field name
}

type responsUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

var user internal.User

var ResponseUser = responsUser{
	int(user.ID),
	user.Email,
	user.Token,
}
