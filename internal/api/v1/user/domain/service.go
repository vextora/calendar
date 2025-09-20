package domain

import "calendarapi/internal/api/v1/user/dto"

type Service interface {
	Register(input *dto.RegisterRequest) (*User, error)
	Login(input *dto.LoginRequest) (string, error)
	FindByID(id int) (*User, error)
}
