package seeder

import (
	"oncomapi/internal/api/v1/user"
	"oncomapi/internal/api/v1/user/dto"
	logs "oncomapi/pkg/logutil"

	"gorm.io/gorm"
)

func seedUser(db *gorm.DB) {
	var count int64

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)

	db.Model([]user.User{}).Count(&count)

	if count > 0 {
		logs.Info("Data exist, skip seeding")
		return
	}

	users := []dto.RegisterRequest{
		{
			Username: "admin",
			Email:    "admin@oncom.com",
			Password: "admin123",
			Detail: &dto.UserDetailRequest{
				FullName: "Admin Oncom",
				Phone:    "081234567890",
				Address:  "Jakarta",
			},
		},
	}

	for _, data := range users {
		_, err := userService.Register(&data)
		if err != nil {
			logs.Error("Failed seed user : %v\n", err)
		}
	}
}
