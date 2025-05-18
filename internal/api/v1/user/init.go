package user

import (
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *Handler {
	repo := NewUserRepository(db)
	service := NewUserService(repo)

	return NewUserHandler(service)
}
