package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/YAWAL/HumanResourceMicroservice/src/database"
)

type Config struct {
	Host     string          `json:"host"`
	Database database.Config `json:"database"`
}

func ReadConfig(filename string) (conf Config, err error) {
	confData, err := ioutil.ReadFile(filename)
	if err != nil {
		return conf, err
	}
	if err = json.Unmarshal(confData, &conf); err != nil {
		return conf, err
	}
	return conf, nil
}
