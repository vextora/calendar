package {{.EntityLower}}

import (
	"gorm.io/gorm"
)

type {{.EntityLower}}Repository struct {
	db *gorm.DB
}

func New{{.Entity}}Repository(db *gorm.DB) Repository {
	return &{{.EntityLower}}Repository{db}
}
