package ratelimiter

import "testing"

func TestNoopRatelimiter(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := NewNoopRateLimiter()
		result1 := <-r.Throttle("a")
		if result1.Error != nil || result1.Delay != 0 {
			t.Error("Unexpected blocking")
		}
		result2 := <-r.Throttle("b")
		if result2.Error != nil || result2.Delay != 0 {
			t.Error("Unexpected blocking")
		}
		result3 := <-r.Throttle("b")
		if result3.Error != nil || result3.Delay != 0 {
			t.Error("Unexpected blocking")
		}
	})
}
