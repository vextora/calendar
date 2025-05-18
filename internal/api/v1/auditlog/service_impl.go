package auditlog

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type auditlogService struct {
	repo Repository
}

func NewAuditlogService(repo Repository) Service {
	return &auditlogService{repo: repo}
}

func (s *auditlogService) Log(ctx context.Context, userID int, action, entity, entityID string, detail json.RawMessage) error {
	log := &AuditLog{
		ID:        uuid.New(),
		UserID:    userID,
		Action:    action,
		Entity:    entity,
		EntityID:  entityID,
		Timestamp: time.Now(),
		Details:   detail,
	}

	return s.repo.Create(ctx, log)
}
