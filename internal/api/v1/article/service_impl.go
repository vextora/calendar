package article

import (
	"fmt"
	"oncomapi/internal/api/v1/article/dto"
	apperror "oncomapi/pkg/errors"
	"oncomapi/pkg/utils"
)

type articleService struct {
	repo Repository
}

func NewArticleService(repo Repository) Service {
	return &articleService{repo}
}

func (s *articleService) GetAll() ([]*Article, error) {
	return s.repo.GetAll()
}

func (s *articleService) GetByID(id uint) (Article, *apperror.AppError) {
	article, err := s.repo.GetByID(id)
	if err != nil {
		return article, err
	}
	return article, nil
}

func (s *articleService) Create(input *dto.CreateRequest) (*Article, error) {
	slug, err := s.GenerateUniqueSlug(input.Title)
	if err != nil {
		return nil, err
	}
	article := &Article{
		Title:   input.Title,
		Content: input.Content,
		Slug:    slug,
	}

	result, err := s.repo.Create(article)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *articleService) Update(input *dto.UpdateRequest) (*Article, error) {
	slug, err := s.GenerateUniqueSlug(input.Title)
	if err != nil {
		return nil, err
	}
	article := &Article{
		ID:      input.ID,
		Title:   input.Title,
		Content: input.Content,
		Slug:    slug,
	}

	result, err := s.repo.Update(article)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *articleService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *articleService) GenerateUniqueSlug(title string) (string, error) {
	baseSlug := utils.GenerateSlug(title)
	finalSlug := baseSlug
	suffix := 1

	for {
		exists, err := s.repo.CheckSlugExists(finalSlug)
		if err != nil {
			return "", err
		}
		if !exists {
			break
		}
		finalSlug = fmt.Sprintf("%s-%d", baseSlug, suffix)
		suffix++
	}
	return finalSlug, nil
}

func (s *articleService) FindByUserID(userID int) ([]*Article, error) {
	article, err := s.repo.FindByUserID(userID)
	if err != nil {
		return article, err
	}

	return article, nil
}
