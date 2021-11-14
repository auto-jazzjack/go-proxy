package Channel

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Channel struct {
	client *http.Client
	port   int
	host   string
}

var headers = map[string]interface{}{
	"Content-Length": nil,
	"Content-Type":   nil,
	"Date":           nil,
}

func NewChannel(host string) *Channel {

	var v = strings.Split(host, ":")

	if len(v) != 2 {
		panic(v)
	}
	var _port, err = strconv.Atoi(v[1])

	if err != nil {
		panic(v)
	}

	return &Channel{
		client: &http.Client{},
		host:   v[0],
		port:   _port,
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

		var body, err2 = ioutil.ReadAll(resp.Body)
		if err2 != nil {
			panic(err2)
		}

		(*res).WriteHeader(resp.StatusCode)
		for k := range headers {
			(*res).Header().Add(k, resp.Header.Get(k))
		}
		(*res).Write(body)
	}
}
