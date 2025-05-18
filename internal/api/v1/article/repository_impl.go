package article

import (
	apperror "oncomapi/pkg/errors"

	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) Repository {
	return &articleRepository{db}
}

func (r *articleRepository) GetAll() ([]*Article, error) {
	var articles []*Article
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *articleRepository) GetByID(id uint) (Article, *apperror.AppError) {
	var article Article
	err := r.db.First(&article, id).Error
	if err != nil {
		return article, apperror.HandleNotFoundError(err, "Article", id)
	}
	return article, nil
}

func (r *articleRepository) Create(article *Article) (*Article, error) {
	if err := r.db.Create(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (r *articleRepository) Update(article *Article) (*Article, error) {
	if err := r.db.Save(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (r *articleRepository) Delete(id uint) error {
	return r.db.Delete(&Article{}, id).Error
}

func (r *articleRepository) CheckSlugExists(slug string) (bool, error) {
	var count int64
	err := r.db.Model(&Article{}).Where("slug = ?", slug).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *articleRepository) FindByUserID(userID int) ([]*Article, error) {
	var articles []*Article
	err := r.db.Where("user_id = ?", userID).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}
