package config

import (
	logs "calendarapi/pkg/logutil"
	"strconv"

	"github.com/spf13/viper"
)

func GetEnvString(key string) string {
	val := viper.GetString(key)
	if val == "" {
		logs.Error("Required env %s not set", key)
	}
	return val
}

func GetEnvInt(key string) int {
	val := viper.GetString(key)
	if val == "" {
		logs.Error("Required env %s not set", key)
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		logs.Error("Invalid int value for %s: %v", key, err)
	}
	return i
}

func GetEnvBool(key string) bool {
	val := viper.GetString(key)
	if val == "" {
		logs.Error("Required env %s not set", key)
	}

	b, err := strconv.ParseBool(val)
	if err != nil {
		logs.Error("Invalid bool value for %s: %v", key, err)
	}
	return b
}
