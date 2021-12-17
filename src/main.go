package main

import (
	"fmt"
	"net/http"
	ad "proxy/src/admin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Channel struct {
	client *http.Client
}

func main() {
	fmt.Println("started")
	var admin = ad.NewAdmin()

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":9393", admin.GetProxy().GetEventLoop())

}
