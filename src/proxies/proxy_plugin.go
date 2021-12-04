package Proxies

import "proxy/proto/go/proxy/config"

type ProxyPlugin interface {
	CanSupport(config.Proxy) bool
}
