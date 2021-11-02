package Channel

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Channel struct {
	client *http.Client
	port   int
}

var headers = map[string]interface{}{
	"Content-Length": nil,
	"Content-Type":   nil,
	"Date":           nil,
}

func NewChannel(port int) *Channel {
	return &Channel{
		client: &http.Client{},
		port:   port,
	}
}
func createRequestForUpstream(req *http.Request, port int) *http.Request {
	var host = strings.Split(req.Host, ":")
	var newHost = "http://"

	if len(host) != 2 {
		panic(newHost)
	} else {
		newHost += host[0]          //host
		newHost += ":"              //port
		newHost += fmt.Sprint(port) //port
		newHost += req.URL.Path     //path
		newHost += req.URL.RawQuery
	}

	var retv, _ = http.NewRequest(req.Method, newHost, req.Body)
	return retv
}

func (h Channel) CallRemote(res *http.ResponseWriter, req *http.Request) {
	var replaced = createRequestForUpstream(req, h.port)
	var resp, err = h.client.Do(replaced)

	if err != nil {
		panic(err)
	} else {
		(*res).WriteHeader(resp.StatusCode)

		var body, err2 = ioutil.ReadAll(resp.Body)
		if err2 != nil {
			panic(err2)
		}

		for k := range headers {
			(*res).Header().Add(k, resp.Header.Get(k))
		}
		(*res).Write(body)
	}
}
