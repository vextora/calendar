package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string             `json:"username" validate:"required"`
	Email    string             `json:"email" validate:"required,email"`
	Password string             `json:"password" validate:"required"`
	Detail   *UserDetailRequest `json:"detail" validate:"required"`
}

type UserDetailRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
}
