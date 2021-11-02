package proxies

import (
	el "proxy/event_loop"
)

type Proxies struct {
	event_loop *el.Eventloop
}

func NewProxies(port int) *Proxies {
	return &Proxies{}
}

func (px *Proxies) GetEventLoop() *el.Eventloop {
	return px.event_loop
}
