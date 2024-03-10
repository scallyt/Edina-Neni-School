package models

type LoginData struct {
	UserName string `json:"username"` // Capitalized field name
	Email    string `json:"email"`    // Capitalized field name
	Password string `json:"password"` // Capitalized field name
}

type ResponseUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
