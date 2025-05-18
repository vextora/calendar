package migration

import (
	"oncomapi/internal/api/v1/article"
	"oncomapi/internal/api/v1/auditlog"
	"oncomapi/internal/api/v1/user"
	logs "oncomapi/pkg/logutil"
	"reflect"

	"gorm.io/gorm"
)

func MigratePostgres(db *gorm.DB) {
	tables := []interface{}{
		&article.Article{},
		&user.User{},
		&user.UserDetail{},
		&auditlog.AuditLog{},
	}

	for _, table := range tables {
		tableName := reflect.TypeOf(table).Elem().Name()
		if err := db.AutoMigrate(table); err != nil {
			logs.Error("Migration failed for table %v: %v", tableName, err)
		} else {
			logs.Info("Migration successful for table %v", tableName)
		}
	}
}
