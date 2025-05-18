package ratelimiter

import (
	"golang.org/x/time/rate"
)

type Limiter struct {
	store LimiterStore
}

func NewLimiter(store LimiterStore) *Limiter {
	return &Limiter{store: store}
}

func (l *Limiter) GetLimiter(key string) *rate.Limiter {
	return l.store.GetLimiter(key)
}
