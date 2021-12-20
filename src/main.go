package main

import (
	"fmt"
	"net/http"
	ad "proxy/src/admin"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Channel struct {
	client *http.Client
}

var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)
	fmt.Println("started")
	var admin = ad.NewAdmin()

	//admin with other thread
	go func() {
		defer wg.Done()
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9001", nil)
	}()

	//proxy go
	go func() {
		defer wg.Done()
		http.ListenAndServe(":9393", admin.GetProxy().GetEventLoop())
	}()

	wg.Wait()
}
