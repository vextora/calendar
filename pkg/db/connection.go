package db

import (
	"gorm.io/gorm"
)

type DBConnections struct {
	Postgres *gorm.DB
}

func InitDB() (*DBConnections, error) {
	postgresDb := InitPostgres()

	return &DBConnections{
		Postgres: postgresDb,
	}, nil
}
