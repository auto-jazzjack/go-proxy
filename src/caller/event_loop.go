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
	handlers map[string]http.Handler
}

func NewEventLoop(cfg *config.Proxy) *Eventloop {
	return &Eventloop{
		channels: []*ch.Channel{},
		rl:       NewRateLimiter(cfg.RateLimit),
		handlers: make(map[string]http.Handler),
	}
}

func (el *Eventloop) RegisterChannel(host string) {
	el.channels = append(el.channels, ch.NewChannel(host))
	el.pos = 0
	el.size = int64(len(el.channels))
}

func (el *Eventloop) RegisterHandler(key string, value http.Handler) {
	if el.handlers[key] != nil {
		panic("Already registered")
	}
	el.handlers[key] = value
}

func (el *Eventloop) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	//start time
	now := time.Now()
	var code = 0

	var adminHandler = el.handlers[req.RequestURI]
	if adminHandler != nil {
		adminHandler.ServeHTTP(res, req)
		return
	}

	if el.rl != nil {
		if el.rl.TryConsume(1) {
			code = el.channels[el.pos].CallRemote(&res, req)
		} else {
			code = el.channels[el.pos].CallTooManyRequest(&res)
		}
	} else {
		code = el.channels[el.pos].CallRemote(&res, req)
	}

	metrics.MeasureCountAndLatency(code, req.RequestURI, now)
}
