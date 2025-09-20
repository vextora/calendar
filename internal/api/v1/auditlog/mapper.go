package auditlog

import "calendarapi/internal/api/v1/auditlog/dto"

func AuditlogResponse(data *AuditLog) dto.AuditlogResponse {
	return dto.AuditlogResponse{
		//ID: uint(data.ID[]),
	}
}
