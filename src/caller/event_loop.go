package Caller

import (
	"net/http"
	"proxy/proto/go/proxy/config"
	ch "proxy/src/channel"
	metrics "proxy/src/metrics"
	"time"
)

type Eventloop struct {
	channels []*ch.Channel
	rl       *RateLimiter
	pos      int64
	size     int64
}

func NewEventLoop(cfg *config.Proxy) *Eventloop {
	return &Eventloop{
		channels: []*ch.Channel{},
		rl:       NewRateLimiter(cfg.RateLimit),
	}
}

func (el *Eventloop) RegisterChannel(host string) {
	el.channels = append(el.channels, ch.NewChannel(host))
	el.pos = 0
	el.size = int64(len(el.channels))
}

func (el *Eventloop) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	//start time
	now := time.Now().UnixMilli()
	var code = 0

	if el.rl != nil {
		if el.rl.TryConsume(1) {
			code = el.channels[el.pos].CallRemote(&res, req)
		} else {
			code = el.channels[el.pos].CallTooManyRequest(&res)
		}
	} else {
		code = el.channels[el.pos].CallRemote(&res, req)
	}

	metrics.MeasureCountAndLatency(code, now)
}
