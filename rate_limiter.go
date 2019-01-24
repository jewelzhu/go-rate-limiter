package go_rate_limiter


import (
	"math"
	"time"
)

// I cannot find a proper Golang rate limiter lib so I wrote one which is referenced from a lua implementation
// https://github.com/openresty/lua-resty-limit-traffic/blob/master/lib/resty/limit/req.lua
// I want requests to be accepted at once if not exceeds limit, delay time returned if it's in the burst, rejected at once if exceeds burst
type RateLimiter struct {
	// how many requests can be accepted withon one second
	rate float64
	// how many burst requests can be allowed
	burst float64

	excess float64
	last time.Time
}

func NewRateLimiter(rate float64, burst float64) *RateLimiter{
	return &RateLimiter{
		rate : rate * 1000,
		burst : burst * 1000,
	}
}

// ok=true and delay=0 : the request can be accepted at once
// ok=true and delay=0 : the request can be executed after the delay
// ok=false : reject the request at once
func (r *RateLimiter) incoming() (delay float64, ok bool) {
	now := time.Now()
	elapsed := now.Sub(r.last) / 1000000   //微秒
	r.excess = math.Max(r.excess - r.rate * math.Abs(float64(elapsed)) / 1000 + 1000, 0)
	if r.excess > r.burst {
		return 0, false
	}
	r.last = now
	return r.excess/r.rate, true
}
