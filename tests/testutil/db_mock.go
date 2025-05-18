package testutil_test

import (
	"oncomapi/internal/api/v1/article/domain"
	apperror "oncomapi/pkg/errors"
)

type MockArticleRepo struct {
	Article *domain.Article
}

func NewMockArticleRepo() *MockArticleRepo {
	return &MockArticleRepo{}
}

func (m *MockArticleRepo) GetByID(id uint) (domain.Article, *apperror.AppError) {
	if m.Article != nil && m.Article.ID == id {
		return *m.Article, nil
	}
	return domain.Article{}, apperror.NotFound("Article not found", id)
}
