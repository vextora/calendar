package seeder

import (
	"oncomapi/internal/api/v1/article"
	logs "oncomapi/pkg/logutil"

	"gorm.io/gorm"
)

func seedArticles(db *gorm.DB) {
	var count int64

	articles := []article.Article{
		{UserID: 1, Title: "Judul 1", Content: "Content 1", Slug: "judul-1"},
		{UserID: 1, Title: "Judul 2", Content: "Content 2", Slug: "judul-2"},
	}

	db.Model([]article.Article{}).Count(&count)

	if count == 0 {
		for _, data := range articles {
			if err := db.Create(&data).Error; err != nil {
				logs.Error("Gagal seeding article : %v\n", err)
			}
		}
	}
}
