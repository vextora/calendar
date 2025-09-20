package seeder

import "gorm.io/gorm"

func SeedAll(db *gorm.DB) {
	seedUser(db)
	seedCalendars(db)
	seedEvents(db)
}
