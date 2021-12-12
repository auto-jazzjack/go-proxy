package Proxies

import (
	"fmt"
	config "proxy/proto/go/proxy/config"
	el "proxy/src/caller"
	wt "proxy/src/watch"
)

type Proxies struct {
	event_loop  *el.Eventloop
	admin_watch chan wt.Event
}

func NewProxies(cfg *config.Proxy, ch chan wt.Event) *Proxies {
	var retv = &Proxies{
		event_loop:  el.NewEventLoop(cfg),
		admin_watch: ch,
	}

	for _, host := range cfg.GetUpstreams() {
		retv.GetEventLoop().RegisterChannel(host)
	}
	go retv.Update()
	return retv
}

func (px *Proxies) GetEventLoop() *el.Eventloop {
	return px.event_loop
}

func (px *Proxies) Update() {
	var evnt = <-px.admin_watch
	fmt.Printf(string(evnt))
}
