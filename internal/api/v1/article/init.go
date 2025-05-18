package article

import (
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *Handler {
	repo := NewArticleRepository(db)
	service := NewArticleService(repo)

	return NewHandler(service)
}
