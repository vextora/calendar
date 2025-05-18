package auditlog

import "gorm.io/gorm"

type AuditLogHandler struct {
	Service Service
}

func Init(db *gorm.DB) *AuditLogHandler {
	repo := NewAuditlogRepository(db)
	service := NewAuditlogService(repo)

	return &AuditLogHandler{
		Service: service,
	}
}
