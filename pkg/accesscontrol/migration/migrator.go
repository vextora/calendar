package migration

import "gorm.io/gorm"

func init() {
	RegisterMigration(Migrate)
}

var migrationFuncs []func(db *gorm.DB) error

func RegisterMigration(fn func(db *gorm.DB) error) {
	migrationFuncs = append(migrationFuncs, fn)
}

func RunAllMigrations(db *gorm.DB) error {
	for _, fn := range migrationFuncs {
		if err := fn(db); err != nil {
			return err
		}
	}

	return nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&CasbinRule{},
		&Role{},
		&Permission{},
		&RolePermission{},
	)
}
