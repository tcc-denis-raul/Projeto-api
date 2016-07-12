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

type Collections struct {
	LanguageIngles string `json:"language_ingles"`
}

type Configuration struct {
	Database    DatabaseType `json:"database"`
	Collections Collections  `json:"collections"`
}

func Conf() (Configuration, error) {
	file, e := ioutil.ReadFile(fmt.Sprintf("%s/paloma.json", DefaultURLConf))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return Configuration{}, e
	}
	var conf Configuration
	json.Unmarshal(file, &conf)
	return conf, nil
}
