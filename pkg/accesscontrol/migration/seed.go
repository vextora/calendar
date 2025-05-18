package migration

import "gorm.io/gorm"

func init() {
	RegisterSeeder(Seed)
}

var seederFuncs []func(db *gorm.DB) error

func RegisterSeeder(fn func(db *gorm.DB) error) {
	seederFuncs = append(seederFuncs, fn)
}

func RunAllSeeders(db *gorm.DB) error {
	for _, fn := range seederFuncs {
		if err := fn(db); err != nil {
			return err
		}
	}

	return nil
}

func Seed(db *gorm.DB) error {
	roles := []Role{
		{Name: "admin", Description: "Administrator"},
		{Name: "editor", Description: "Editor"},
		{Name: "user", Description: "Regular User"},
	}
	for _, r := range roles {
		db.FirstOrCreate(&r, Role{Name: r.Name})
	}

	permissions := []Permission{
		{Action: "read", Resource: "article"},
		{Action: "write", Resource: "article"},
		{Action: "delete", Resource: "article"},
	}

	for _, p := range permissions {
		db.FirstOrCreate(&p, Permission{Action: p.Action, Resource: p.Resource})
	}

	var admin Role
	db.First(&admin, "name = ?", "admin")

	var perms []Permission
	db.Find(&perms)

	for _, p := range perms {
		db.FirstOrCreate(&RolePermission{}, RolePermission{
			RoleID:       admin.ID,
			PermissionID: p.ID,
		})
	}

	return nil
}
