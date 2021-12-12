package Repository

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/golang/protobuf/jsonpb"

	config "proxy/proto/go/proxy/config"

	yamlToJson "github.com/ghodss/yaml"
)

type Repository struct {
	version string
}

type RepositoryImpl interface {
	GetConf()
}

func NewRepository() *Repository {
	return &Repository{version: getLatestVersion("../configuration")}
}

func (rp *Repository) GetConf() *config.Proxy {

	yamlFile, err := ioutil.ReadFile("../configuration/proxy" + rp.version + ".yaml")
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

func getLatestVersion(path string) string {
	var files, err = ioutil.ReadDir(path)

	if err != nil {
		fmt.Printf(err.Error())
		panic("Cannot get latest file")
	}

	if len(files) <= 0 {
		panic("Cannot get any file")
	}

	var retv = files[0].Name()

	for _, itr := range files {
		if strings.Count(itr.Name(), ".") != 1 {
			continue
		}
		if retv < itr.Name() {
			retv = itr.Name()
		}
	}

	retv = strings.Split(retv, ".")[0]           //select head
	return strings.Replace(retv, "proxy", "", 1) //remove common part
}

func (rp *Repository) CreateRevision() {
	//copy current proto file with version
}
