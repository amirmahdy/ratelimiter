package ratelimit

import "time"

type TimeProvider interface {
	Now() time.Time
}
