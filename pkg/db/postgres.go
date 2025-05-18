package db

import (
	"fmt"
	"oncomapi/pkg/config"
	logs "oncomapi/pkg/logutil"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() *gorm.DB {
	host := config.GetEnvString(config.PostgreDbHost)
	port := config.GetEnvString(config.PostgreDbPort)
	user := config.GetEnvString(config.PostgreDbUser)
	pswd := config.GetEnvString(config.PostgreDbPassword)
	dbname := config.GetEnvString(config.PostgreDbName)
	ssl := config.GetEnvString(config.PostgreDbSsl)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		host, user, pswd, dbname, port, ssl,
	)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Error("Error connecting to database : %v", err)
	}

	logs.Info("Database connected successfully")

	return db
}
