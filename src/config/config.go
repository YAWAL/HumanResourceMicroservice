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

func ReadConfig(path string) (conf *Config, err error) {
	confData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(confData, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
