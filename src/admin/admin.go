package Admin

import (
	"net/http"
	Proxies "proxy/src/proxies"
	px "proxy/src/proxies"
	rp "proxy/src/repository"
	wt "proxy/src/watch"
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

func (adm *Admin) Update(data string) {
	var err = adm.repo.CreateRevision(data)
	if err != nil {
		panic("Can not create revision")
	}
}

func (adm *Admin) GetProxy() *px.Proxies {
	return adm.proxy
}

func (adm *Admin) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "/update/configuration" {

		len := req.ContentLength
		body := make([]byte, len)
		req.Body.Read(body)

		adm.repo.CreateRevision(string(body))
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Ok"))
		return
	}

	res.WriteHeader(http.StatusNotFound)
}
