package main

import (
	"fmt"
	"net/http"
	el "proxy/src/event_loop"
)

type Channel struct {
	client *http.Client
}

func main() {
	fmt.Print("started")
	var handler = el.NewEventLoop(9290)
	http.ListenAndServe(":9393", handler)
}
