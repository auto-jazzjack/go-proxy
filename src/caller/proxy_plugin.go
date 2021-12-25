package Caller

import (
	"net/http"
	"proxy/proto/go/proxy/config"
)

type ProxyPluginImpl interface {
	CanSupport(*config.Proxy) bool
	TryConsume(*http.Request) bool
	IsEnabled() bool
	Order() int
	FallbackHttpStatus() int
}

type ProxyPlugin struct {
}
