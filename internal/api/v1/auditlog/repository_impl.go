package auditlog

import (
	"context"

	"gorm.io/gorm"
)

type auditlogRepository struct {
	db *gorm.DB
}

func NewAuditlogRepository(db *gorm.DB) Repository {
	return &auditlogRepository{db}
}

func (r *auditlogRepository) Create(ctx context.Context, log *AuditLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}
