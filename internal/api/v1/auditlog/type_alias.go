package auditlog

import "oncomapi/internal/api/v1/auditlog/domain"

type (
	AuditLog   = domain.AuditLog
	Repository = domain.AuditLogRepository
	Service    = domain.AuditLogService
)
