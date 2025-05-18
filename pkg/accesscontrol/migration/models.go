package migration

import "time"

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	PType string `gorm:"size:100;index"`
	V0    string `gorm:"size:100;index"`
	V1    string `gorm:"size:100;index"`
	V2    string `gorm:"size:100;index"`
	V3    string `gorm:"size:100"`
	V4    string `grom:"size:100"`
	V5    string `gorm:"size:100"`
}

type Role struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"uniqueIndex;size:100"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Permission struct {
	ID          uint   `gorm:"primaryKey"`
	Action      string `gorm:"size:50"`
	Resource    string `gorm:"size:100"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Roles       []Role `gorm:"many2many:role_permissions;"`
}

type RolePermission struct {
	RoleID       uint
	PermissionID uint
}

func (CasbinRule) TableName() string {
	return "casbin_rules"
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
