package dto

type RequestLogin struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RequestJoin struct {
	Username string `json:"username" validate:"required,min=4,max=20"`
	Password string `json:"password" validate:"required,min=4,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone,omitempty" validate:"omitempty"`
}
