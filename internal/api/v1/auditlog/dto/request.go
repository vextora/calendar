package dto

type AuditlogRequest struct {
	ID uint `json:"id" validate:"required"`
}
