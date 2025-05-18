package domain

import (
	"context"
	"encoding/json"
)

type AuditLogService interface {
	Log(ctx context.Context, userID int, action, entity, entityID string, details json.RawMessage) error
}
