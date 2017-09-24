package ratelimiter

// NoopRatelimiter implements Throttler without ever blocking
type NoopRatelimiter struct{}

// Throttle immediately returns an empty result
func (l *NoopRatelimiter) Throttle(identifier string) <-chan Result {
	out := make(chan Result)
	go func() {
		out <- Result{}
		close(out)
	}()
	return out
}

// NewNoopRateLimiter returns a Throttler that never blocks
func NewNoopRateLimiter() Throttler {
	return &NoopRatelimiter{}
}
