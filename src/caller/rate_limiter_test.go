package Caller

import (
	"fmt"
	"proxy/proto/go/proxy/config"
	"testing"
)

func TestFoo(t *testing.T) {
	// todo test code
	var rate = NewRateLimiter(&config.Proxy{})

	if v, ok := interface{}(rate).(ProxyPluginImpl); ok {
		fmt.Println(v, ok)
	}

}
