package domain

import apperror "oncomapi/pkg/errors"

type Repository interface {
	GetAll() ([]*Article, error)
	GetByID(id uint) (Article, *apperror.AppError)
	Create(article *Article) (*Article, error)
	Update(article *Article) (*Article, error)
	Delete(id uint) error

	FindByUserID(userID int) ([]*Article, error)

	CheckSlugExists(slug string) (bool, error)
}
