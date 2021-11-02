package Proxies

import (
	"fmt"
	"io/ioutil"
	"log"
	config "proxy/proto/go/proxy/config"

	yaml "gopkg.in/yaml.v2"
)

func GetConf() *config.Proxy {

	yamlFile, err := ioutil.ReadFile("../configuration/proxy.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var v = &config.Proxy{}
	err = yaml.Unmarshal(yamlFile, v)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(v)
	return v
}
