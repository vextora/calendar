package {{.EntityLower}}

import "gorm.io/gorm"

func Init(db *gorm.DB) *{{.Entity}}Handler {
	repo := New{{.Entity}}Repository(db)
	service := New{{.Entity}}Service(repo)

	return New{{.Entity}}Handler(service)
}
