package oncom

import (
	"gorm.io/gorm"
)

type oncomRepository struct {
	db *gorm.DB
}

func NewOncomRepository(db *gorm.DB) Repository {
	return &oncomRepository{db}
}
