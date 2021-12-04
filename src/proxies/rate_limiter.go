package Proxies

import (
	"proxy/proto/go/proxy/config"
)

type RateLimiterImpl ProxyPlugin
type RateLimiter struct {
	qps             int64
	threshhold      int64
	threshholdRatio float64
	limiter         rm.ratelimit
}

func NewRateLimiter(cf config.RateLimit) *RateLimiter {

	if !canSupport(cf) {
		return nil
	}

	return &RateLimiter{}
}

func (rl *RateLimiter) CanSupport(cf config.Proxy) bool {
	return canSupport(*cf.RateLimit)
}

func canSupport(cf config.RateLimit) bool {
	if cf.GetQps() <= 0 || cf.GetThreshholdRatio() < 0 || cf.GetThreshhold() < 0 {
		return false
	}

	return true
}

func (rl *RateLimiter) TryConsume(count int64) bool {

}
