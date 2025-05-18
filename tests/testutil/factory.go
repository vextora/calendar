package testutil_test

import "oncomapi/internal/api/v1/article/domain"

func FakeArticle() *domain.Article {
	return &domain.Article{
		ID:      1,
		Title:   "Dummy Title",
		Content: "Dummy content",
	}
}
