package domain

import (
	"oncomapi/internal/api/v1/article/dto"
	apperror "oncomapi/pkg/errors"
)

type Service interface {
	GetAll() ([]*Article, error)
	GetByID(id uint) (Article, *apperror.AppError)
	Create(input *dto.CreateRequest) (*Article, error)
	Update(input *dto.UpdateRequest) (*Article, error)
	Delete(id uint) error

	FindByUserID(userID int) ([]*Article, error)

	GenerateUniqueSlug(title string) (string, error)
}
