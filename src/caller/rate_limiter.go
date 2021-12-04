package Caller

import (
	"proxy/proto/go/proxy/config"
	"time"

	rm "golang.org/x/time/rate"
)

type RateLimiterImpl ProxyPlugin
type RateLimiter struct {
	qps             int64
	threshhold      int64
	threshholdRatio float64
	limiter         rm.Limiter
}

func NewRateLimiter(cf *config.RateLimit) *RateLimiter {

	if !canSupport(cf) {
		return nil
	}

	return &RateLimiter{
		qps:             cf.GetQps(),
		threshhold:      cf.GetThreshhold(),
		threshholdRatio: cf.GetThreshholdRatio(),
		limiter:         *rm.NewLimiter(rm.Limit(cf.GetQps()), int(cf.GetQps())),
	}
}

func (rl *RateLimiter) CanSupport(cf *config.Proxy) bool {
	return canSupport(cf.RateLimit)
}

func canSupport(cf *config.RateLimit) bool {
	if cf.GetQps() <= 0 || cf.GetThreshholdRatio() < 0 || cf.GetThreshhold() < 0 || cf.GetThreshholdRatio() > 1 {
		return false
	}

	return true
}

func (rl *RateLimiter) TryConsume(count int64) bool {
	return rl.limiter.AllowN(time.Now(), int(count))
}
