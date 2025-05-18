package oncom

import "gorm.io/gorm"

func Init(db *gorm.DB) *OncomHandler {
	repo := NewOncomRepository(db)
	service := NewOncomService(repo)

	return NewOncomHandler(service)
}
