package main

import (
	"fmt"
	"net/http"
	px "proxy/src/proxies"
)

type Channel struct {
	client *http.Client
}

func main() {
	fmt.Print("started")
	var pro = px.NewProxies(px.GetConf())
	http.ListenAndServe(":9393", pro.GetEventLoop())
}
