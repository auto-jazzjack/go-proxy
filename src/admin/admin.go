package Admin

import (
	px "proxy/src/proxies"
	wt "proxy/src/watch"
)

type Admin struct {
	proxy      *px.Proxies
	connection chan wt.Event
}

type AdminImpl interface {
	update()
}

func NewAdmin(pxx *px.Proxies) *Admin {

	var _chan = make(chan wt.Event)

	pxx.SetAdminWatch(_chan)

	return &Admin{
		pxx, _chan,
	}
}

func (adm *Admin) update() {

}
