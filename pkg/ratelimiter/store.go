package ratelimiter

import "golang.org/x/time/rate"

type LimiterStore interface {
	GetLimiter(key string) *rate.Limiter
}
