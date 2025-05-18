package ratelimiter

func InitDefaultLimiter() *Limiter {
	cfg := LoadConfig()
	store := NewMemoryStore(cfg)
	return NewLimiter(store)
}
