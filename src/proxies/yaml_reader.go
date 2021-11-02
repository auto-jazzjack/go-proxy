package Proxies

import (
	"io/ioutil"
	"log"
	proto "github.com/golang/protobuf/proto"
	"proxy/proto/config"
   
)

func getConf(cfg config) {

	proto.
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
