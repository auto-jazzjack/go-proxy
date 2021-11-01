package main

import (
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(rew http.ResponseWriter, req *http.Request) {
	rew.WriteHeader(200)
	rew.Write([]byte("helloWorld"))
}
func main() {
	var v = HttpHandler{}
	http.ListenAndServe(":9393", v)
}
