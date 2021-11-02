package Event_loop

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
		channels: []*ch.Channel{},
	}
}

func (el *Eventloop) RegisterChannel(host string) {
	el.channels = append(el.channels, ch.NewChannel0(host))
	el.pos = 0
	el.size = int64(len(el.channels))
}

func (el *Eventloop) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	el.channels[el.pos].CallRemote(&res, req)
}
