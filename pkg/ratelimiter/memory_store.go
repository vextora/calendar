package ratelimiter

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type MemoryStore struct {
	clients    map[string]*rate.Limiter
	mu         sync.Mutex
	rate       rate.Limit
	burst      int
	expiration time.Duration
}

func NewMemoryStore(cfg Config) *MemoryStore {
	limit := rate.Every(time.Duration(cfg.ReqPerSecond) * time.Second)

	return &MemoryStore{
		clients:    make(map[string]*rate.Limiter),
		rate:       limit,
		burst:      cfg.Burst,
		expiration: time.Duration(cfg.TTLSeconds) * time.Second,
	}
}

func (m *MemoryStore) GetLimiter(key string) *rate.Limiter {
	m.mu.Lock()
	defer m.mu.Unlock()

	limiter, exists := m.clients[key]
	if !exists {
		limiter = rate.NewLimiter(m.rate, m.burst)
		m.clients[key] = limiter

		go func() {
			time.Sleep(m.expiration)
			m.mu.Lock()
			defer m.mu.Unlock()
			delete(m.clients, key)
		}()
	}

	return limiter
}
