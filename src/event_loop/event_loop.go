package eventloop

import (
	"net/http"
	ch "proxy/src/channel"
)

type Eventloop struct {
	channels []*ch.Channel
	pos      int64
	size     int64
}

func NewEventLoop() *Eventloop {
	return &Eventloop{
		channels: []*ch.Channel{ch.NewChannel(9290)},
	}
}

func (el Eventloop) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	el.channels[el.pos].CallRemote(&res, req)
}
