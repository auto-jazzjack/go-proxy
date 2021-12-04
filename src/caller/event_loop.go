package Caller

import (
	"net/http"
	"proxy/proto/go/proxy/config"
	ch "proxy/src/channel"
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
	if el.rl != nil {
		if el.rl.TryConsume(1) {
			el.channels[el.pos].CallRemote(&res, req)
		} else {
			el.channels[el.pos].CallTooManyRequest(&res)
		}
		return
	}

	el.channels[el.pos].CallRemote(&res, req)
}
