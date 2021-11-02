package Event_loop

import (
	"net/http"
	ch "proxy/channel"
)

type Eventloop struct {
	channels []*ch.Channel
	pos      int64
	size     int64
}

func NewEventLoop(port int) *Eventloop {
	return &Eventloop{
		channels: []*ch.Channel{ch.NewChannel(port)},
	}
}

func (el Eventloop) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	el.channels[el.pos].CallRemote(&res, req)
}
