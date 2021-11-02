package main

import (
	"fmt"
	"net/http"
	el "proxy/event_loop"
)

type Channel struct {
	client *http.Client
}

func main() {
	fmt.Print("started")
	var handler = el.NewEventLoop(3000)
	http.ListenAndServe(":9393", handler)
}
