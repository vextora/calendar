package config

import (
	logs "calendarapi/pkg/logutil"
	"os"

	"github.com/spf13/viper"
)

func InitDotenv() {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	// Load default (.env) => assume as base config production
	viper.SetConfigName(".env")
	if err := viper.MergeInConfig(); err != nil {
		logs.Debug("Warning: .env (default) not found : %s", err)
	}

	// Read ENV from environment variable or .env
	env := os.Getenv("ENV")
	if env == "" {
		env = viper.GetString("ENV")
	}
	if env == "" {
		env = "production"
	}
	logs.Debug("File env yg diload : %s", env)

	// Load override if APP_ENV is not production
	if env != "production" {
		viper.SetConfigName(".env." + env)
		if err := viper.MergeInConfig(); err != nil {
			logs.Debug("Warning: .env.%s not found : %s", env, err)
		}
	}

	logs.Info("Configuration loaded for environment: %s", env)
}
