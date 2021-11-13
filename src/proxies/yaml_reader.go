package Proxies

import (
	"fmt"
	"io/ioutil"
	config "proxy/proto/go/proxy/config"

	yaml "gopkg.in/yaml.v2"
)

func GetConf() *config.Proxy {

	yamlFile, err := ioutil.ReadFile("../configuration/proxy.yaml")
	if err != nil {
		panic("yamlFile.Get err")
	}

	var v = &config.Proxy{}
	err = yaml.Unmarshal(yamlFile, v)
	if err != nil {
		panic("marformed" + string(yamlFile))
	}

	fmt.Println(v)
	return v
}
