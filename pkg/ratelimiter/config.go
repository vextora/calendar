package ratelimiter

import "calendarapi/pkg/config"

type Config struct {
	ReqPerSecond int
	Burst        int
	TTLSeconds   int
}

func LoadConfig() Config {
	return Config{
		ReqPerSecond: config.GetEnvInt(config.RateLimitPerSecond),
		Burst:        config.GetEnvInt(config.RateLimitBurst),
		TTLSeconds:   config.GetEnvInt(config.RateLimitTtlSeconds),
	}
}
