package main

import (
	"fmt"
	"net/http"
	ad "proxy/admin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Channel struct {
	client *http.Client
}

func main() {

	fmt.Println("started")
	var admin = ad.NewAdmin()
	admin.GetProxy().GetEventLoop().RegisterHandler("/metrics", promhttp.Handler())
	http.ListenAndServe(":9393", admin.GetProxy().GetEventLoop())
}
