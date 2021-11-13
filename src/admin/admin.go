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

func NewAdmin() *Admin {
	return &Admin{}
}

func (adm *Admin) update() {

}
