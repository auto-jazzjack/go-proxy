package Caller

import (
	"net/http"
	"proxy/proto/go/proxy/config"
	"time"

	rm "golang.org/x/time/rate"
)

type RateLimiterImpl ProxyPluginImpl
type RateLimiter struct {
	qps             int64
	threshhold      int64
	threshholdRatio float64
	limiter         rm.Limiter
	enabled         bool
}

func NewRateLimiter(cf *config.Proxy) *RateLimiter {

	return &RateLimiter{
		qps:             cf.RateLimit.GetQps(),
		threshhold:      cf.RateLimit.GetThreshhold(),
		threshholdRatio: cf.RateLimit.GetThreshholdRatio(),
		limiter:         *rm.NewLimiter(rm.Limit(cf.RateLimit.GetQps()), int(cf.RateLimit.GetQps())),
		enabled:         true,
	}
}

func (rl *RateLimiter) CanSupport(cf *config.Proxy) bool {
	var ratelimit = cf.RateLimit
	if ratelimit == nil {
		return false
	}
	if ratelimit.GetQps() <= 0 || ratelimit.GetThreshholdRatio() < 0 || ratelimit.GetThreshhold() < 0 || ratelimit.GetThreshholdRatio() > 1 {
		return false
	}
	return true
}

func (rl *RateLimiter) TryConsume(*http.Request) bool {
	return rl.limiter.AllowN(time.Now(), 1)
}

func (rl *RateLimiter) IsEnabled() bool {
	return rl.enabled
}

func (rl *RateLimiter) Order() int {
	return 0
}

func (rl *RateLimiter) FallbackHttpStatus() int {
	return http.StatusTooManyRequests
}
