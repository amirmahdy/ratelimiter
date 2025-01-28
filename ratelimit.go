package ratelimit

import (
	"sync"
	"time"
)

type Token struct {
	mx              sync.Mutex
	token           int // Total Tokens
	refreshRate     int // token they rec each sec
	limit           int // Limit token
	lastRefreshTime time.Time
}

func NewToken(limit int, refreshRate int, tm TimeProvider) *Token {
	return &Token{
		limit:           limit,
		token:           limit,
		lastRefreshTime: tm.Now(),
		refreshRate:     refreshRate,
	}
}

func (tk *Token) RefreshTokens(tm TimeProvider) {
	currentTime := tm.Now()
	newTokens := tk.refreshRate * int(currentTime.Sub(tk.lastRefreshTime).Seconds())
	if newTokens > 0 {
		tk.token = min(tk.limit, newTokens+tk.token)
		tk.lastRefreshTime = tm.Now()
	}
}

func (tk *Token) Allow(tm TimeProvider) bool {
	tk.mx.Lock()
	defer tk.mx.Unlock()

	tk.RefreshTokens(tm)
	if tk.token > 0 {
		tk.token -= 1
		return true
	}
	return false
}
