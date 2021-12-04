package Proxies

import (
	config "proxy/proto/go/proxy/config"
	el "proxy/src/caller"
	wt "proxy/src/watch"
)

type Proxies struct {
	event_loop  *el.Eventloop
	admin_watch chan wt.Event
}

func NewProxies(cfg *config.Proxy) *Proxies {
	var retv = &Proxies{
		event_loop: el.NewEventLoop(cfg),
	}

	for _, host := range cfg.GetUpstreams() {
		retv.GetEventLoop().RegisterChannel(host)
	}
	return retv
}

func (px *Proxies) GetEventLoop() *el.Eventloop {
	return px.event_loop
}
