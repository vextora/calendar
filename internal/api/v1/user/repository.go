package user

import (
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{db}
}

func (repo *userRepository) Save(user *User) (*User, error) {
	err := repo.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepository) FindByID(id int) (*User, error) {
	var user User
	err := repo.db.First(&user, id).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (repo *userRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) CreateDetail(detail *UserDetail) (*UserDetail, error) {
	err := repo.db.Create(detail).Error
	if err != nil {
		return nil, err
	}
	return detail, err
}
