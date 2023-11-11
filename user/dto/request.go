package dto

type RequestLogin struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RequestJoin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
