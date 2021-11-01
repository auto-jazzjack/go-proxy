package Channel

import (
	"fmt"
	"net/http"
	"strings"
)

type Channel struct {
	client *http.Client
	port   int
}

func NewChannel(port int) *Channel {
	return &Channel{
		client: &http.Client{},
		port:   port,
	}
}
func createRequestForUpstream(req *http.Request, port int) *http.Request {
	var host = strings.Split(req.Host, ":")
	var newHost = ""

	if len(host) != 2 {
		panic(newHost)
	} else {
		newHost = host[0] + ":" + fmt.Sprint(port)
	}

	var retv, _ = http.NewRequest(req.Method, newHost, nil)
	return retv
}
func (h Channel) CallRemote(res *http.ResponseWriter, req *http.Request) {
	var replaced = createRequestForUpstream(req, h.port)
	var resp, err = h.client.Do(replaced)

	if err != nil {
		panic(err)
	} else {
		(*res).WriteHeader(resp.StatusCode)
		(*res).Write([]byte("test"))
	}
}
