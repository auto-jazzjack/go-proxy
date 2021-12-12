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

const PREFIX = "proxy"
const POSTFIX = ".yaml"

func NewRepository() *Repository {
	return &Repository{version: getLatestVersion("../configuration")}
}

func (rp *Repository) GetConf() *config.Proxy {

	yamlFile, err := ioutil.ReadFile("../configuration/" + PREFIX + rp.version + POSTFIX)
	if err != nil {
		panic("yamlFile.Get err")
	}

	var json, err1 = yamlToJson.YAMLToJSON(yamlFile)

	if err1 != nil {
		panic("Connot parse yaml to json")
	}

	var v, err2 = parserConfig(json)

	if err2 != nil {
		panic("Connot parse json to protobuf")
	}
	return v.Config
}

func parserConfig(jsonByte []byte) (*config.Config, error) {
	var v = &config.Config{}
	var err = jsonpb.UnmarshalString(string(jsonByte), v)

	if err != nil {
		return nil, err
	}
	return v, nil
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

	retv = strings.Split(retv, ".")[0]          //select head
	return strings.Replace(retv, PREFIX, "", 1) //remove common part
}

func (rp *Repository) CreateRevision(data string) error {

	var _, validate = parserConfig([]byte(data))

	if validate != nil {
		return validate
	}

	var err = ioutil.WriteFile("../configuration/"+PREFIX+rp.version+POSTFIX, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
	//copy current proto file with version
}
