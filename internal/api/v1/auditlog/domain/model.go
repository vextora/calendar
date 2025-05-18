package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID        uuid.UUID       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int             `json:"user_id"`
	Action    string          `json:"action"`
	Entity    string          `json:"entity"`
	EntityID  string          `json:"entity_id"`
	Timestamp time.Time       `json:"timestamp"`
	Details   json.RawMessage `json:"details"`
}
