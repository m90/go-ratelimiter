package ratelimiter

import (
	"time"
)

// GetSetter needs to be implemented by any cache that is
// to be used for storing limits
type GetSetter interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, expiry time.Duration)
}

// Throttler needs to be implemented by any rate limiter
type Throttler interface {
	Throttle(identifier string) <-chan Result
}
