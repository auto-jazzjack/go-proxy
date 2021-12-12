package main

import (
	"fmt"
	"net/http"
	ad "proxy/src/admin"
)

type Channel struct {
	client *http.Client
}

func main() {
	fmt.Println("started")
	var admin = ad.NewAdmin()

	go http.ListenAndServe(":9000", admin)
	http.ListenAndServe(":9393", admin.GetProxy().GetEventLoop())
}
