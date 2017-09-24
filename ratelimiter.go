package ratelimiter

import (
	"errors"
	"time"
)

var (
	errInvalidCache = errors.New("invalid value in cache")
)

// Limiter can be used to rate limit operations
// based on an identifier and a threshold value
type Limiter struct {
	threshold time.Duration
	cache     GetSetter
}

// Result describes the outcome of a `Throttle` call
type Result struct {
	Error error
	Delay time.Duration
}

// Throttle returns a channel that blocks until the configured
// rate limit has been satisfied. The channel will send a `Result` exactly
// once before closing, containing information on the
// applied rate limiting or possible errors that occured
func (l *Limiter) Throttle(identifier string) <-chan Result {
	out := make(chan Result)
	go func() {
		if value, found := l.cache.Get(identifier); found {
			if timeout, ok := value.(time.Time); ok {
				remaining := time.Until(timeout)
				l.cache.Set(
					identifier,
					timeout.Add(l.threshold),
					remaining,
				)
				time.Sleep(remaining)
				out <- Result{Delay: remaining}
			} else {
				out <- Result{Error: errInvalidCache}
			}
		} else {
			l.cache.Set(identifier, time.Now().Add(l.threshold), l.threshold)
			out <- Result{}
		}
		close(out)
	}()
	return out
}

// New creates a new Throttler using Limiter. `threshold` defines the
// enforced minimum distance between two calls of the
// instance's `Throttle` method using the same identifier
func New(threshold time.Duration, cache GetSetter) Throttler {
	return &Limiter{
		cache:     cache,
		threshold: threshold,
	}
}
