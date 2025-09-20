package migration

import (
	"calendarapi/internal/api/v1/calendars"
	"calendarapi/internal/api/v1/events"
	"calendarapi/internal/api/v1/user"
	logs "calendarapi/pkg/logutil"
	"reflect"

	"gorm.io/gorm"
)

func MigratePostgres(db *gorm.DB) {
	tables := []any{
		&user.User{},
		&user.UserDetail{},
		&calendars.Calendars{},
		&events.Events{},
	}

	for _, table := range tables {
		tableName := reflect.TypeOf(table).Elem().Name()

		exists := db.Migrator().HasTable(table)

		if err := db.AutoMigrate(table); err != nil {
			logs.Error("Migration failed for table %v: %v", tableName, err)
			continue
		}

		if exists {
			logs.Warn("No schema change detected for table %v", tableName)
		} else {
			logs.Info("Migration successful for table %v", tableName)
		}
	}
}
