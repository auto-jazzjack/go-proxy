package Caller

import (
	"net/http"
	ch "proxy/channel"
	metrics "proxy/metrics"
	"proxy/proto/go/proxy/config"
	"sort"
	"time"
)

type Eventloop struct {
	channels []*ch.Channel
	plugins  []ProxyPluginImpl
	pos      int64
	size     int64
	timeout  int64
	handlers map[string]http.Handler
}

func NewEventLoop(cfg *config.Proxy) *Eventloop {
	return &Eventloop{
		channels: []*ch.Channel{},
		handlers: make(map[string]http.Handler),
		plugins:  createProxyPlugin(cfg),
		timeout:  cfg.GetTimeout(),
	}
}

func createProxyPlugin(cfg *config.Proxy) []ProxyPluginImpl {

	var retv = []ProxyPluginImpl{}

	if NewRateLimiter(cfg).IsEnabled() {
		retv = append(retv, NewRateLimiter(cfg))
	}

	sort.Slice(retv, func(i, j int) bool {
		return retv[i].Order() < retv[j].Order()
	})

	return retv
}

func (el *Eventloop) RegisterChannel(host string) {

	el.channels = append(el.channels, ch.NewChannel(host, el.timeout))
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

	for idx := range el.plugins {
		if !el.plugins[idx].TryConsume(req) {
			res.WriteHeader(el.plugins[idx].FallbackHttpStatus())
			return
		}
	}

	code = el.channels[el.pos].CallRemote(&res, req)

	metrics.MeasureCountAndLatency(code, req.RequestURI, now)
}
