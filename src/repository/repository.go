package Repository

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/jsonpb"

	config "proxy/proto/go/proxy/config"

	yamlToJson "github.com/ghodss/yaml"
)

type Repository struct {
}

type RepositoryImpl interface {
	GetConf()
}

func NewRepository() *Repository {
	return &Repository{}
}

func (rp *Repository) GetConf() *config.Proxy {

	yamlFile, err := ioutil.ReadFile("../configuration/proxy.yaml")
	if err != nil {
		panic("yamlFile.Get err")
	}

	var v = &config.Config{}

	var json, err1 = yamlToJson.YAMLToJSON(yamlFile)
	jsonpb.UnmarshalString(string(json), v)

	if err1 != nil {
		panic("marformed" + string(yamlFile))
	}

	fmt.Println(v)
	return v.Config
}
