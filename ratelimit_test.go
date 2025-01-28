package ratelimit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type timeProvider struct {
	expectedTime time.Time
}

func (tm timeProvider) Now() time.Time {
	return tm.expectedTime
}

func TestNewToken(t *testing.T) {
	customerMap := make(map[int]*Token)
	user1 := 1
	currentTime, _ := time.Parse("2006-01-02T15:04:05", "2025-01-01T00:00:00")

	tm := timeProvider{expectedTime: currentTime}
	customerMap[user1] = NewToken(3, 1, tm)

	require.Equal(t, 3, customerMap[user1].limit)
	require.Equal(t, 3, customerMap[user1].token)
	require.Equal(t, currentTime, customerMap[user1].lastRefreshTime)
}

func TestRefreshToken(t *testing.T) {
	customerMap := make(map[int]*Token)
	user1 := 1
	currentTime, _ := time.Parse("2006-01-02T15:04:05", "2025-01-01T00:00:00")

	tm := timeProvider{expectedTime: currentTime}
	customerMap[user1] = NewToken(3, 1, tm)

	tm = timeProvider{expectedTime: currentTime.Add(time.Second * 10)}
	customerMap[user1].RefreshTokens(tm)

	require.Equal(t, 3, customerMap[user1].limit)
	require.Equal(t, 3, customerMap[user1].token)
	require.Equal(t, currentTime.Add(time.Second*10), customerMap[user1].lastRefreshTime)
}

func TestRateLimitBlocked(t *testing.T) {
	customerMap := make(map[int]*Token)
	user1 := 1
	currentTime, _ := time.Parse("2006-01-02T15:04:05", "2025-01-01T00:00:00")

	tm := timeProvider{expectedTime: currentTime}

	customerMap[user1] = NewToken(3, 1, tm)

	// Making 3 requests
	for i := 0; i < 3; i++ {
		customerMap[user1].Allow(tm)
	}
	resp := customerMap[user1].Allow(tm) // 4th request
	require.Equal(t, false, resp)
}

func TestRateLimitAllow(t *testing.T) {
	customerMap := make(map[int]*Token)
	user1 := 1
	currentTime, _ := time.Parse("2006-01-02T15:04:05", "2025-01-01T00:00:00")

	tm := timeProvider{expectedTime: currentTime}
	customerMap[user1] = NewToken(3, 1, tm)

	for i := 0; i < 2; i++ {
		customerMap[user1].Allow(tm)
	}
	resp := customerMap[user1].Allow(tm) // 3rd request
	require.Equal(t, true, resp)
}

func TestRateLimitAllowTime(t *testing.T) {
	customerMap := make(map[int]*Token)

	// Defining users
	user1 := 1
	currentTime, _ := time.Parse("2006-01-02T15:04:05", "2025-01-01T00:00:00")

	tm := timeProvider{expectedTime: currentTime}
	customerMap[user1] = NewToken(3, 1, tm)

	for i := 0; i < 3; i++ {
		customerMap[user1].Allow(tm)
	}

	tm = timeProvider{currentTime.Add(time.Millisecond * 400)}

	resp := customerMap[user1].Allow(tm) // 4th request
	require.Equal(t, false, resp)

	tm = timeProvider{currentTime.Add(time.Second * 1)}
	resp = customerMap[user1].Allow(tm) // 5th request
	require.Equal(t, true, resp)
}
