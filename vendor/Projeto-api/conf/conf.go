package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const DefaultURLConf = "/etc"

type DatabaseType struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type Configuration struct {
	Database DatabaseType `json:"database"`
}

func Conf(path_c string) (DatabaseType, error) {
	path := DefaultURLConf
	if path_c != "" {
		path = path_c
	}
	file, e := ioutil.ReadFile(fmt.Sprintf("%s/paloma.json", path))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return DatabaseType{}, e
	}
	var conf Configuration
	json.Unmarshal(file, &conf)
	return conf.Database, nil
}
