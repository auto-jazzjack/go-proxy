package Admin

import (
	Proxies "proxy/proxies"
	px "proxy/proxies"
	rp "proxy/repository"
	wt "proxy/watch"
)

type Admin struct {
	proxy      *px.Proxies
	repo       *rp.Repository
	connection chan wt.Event
}

type AdminImpl interface {
	update()
}

func NewAdmin() *Admin {

	var repo = rp.NewRepository()
	var proxy = Proxies.NewProxies(repo.GetConf())
	var connection = make(chan wt.Event)

	return &Admin{
		proxy,
		repo,
		connection,
	}
}

func (adm *Admin) update() {

}

func (adm *Admin) GetProxy() *px.Proxies {
	return adm.proxy
}
