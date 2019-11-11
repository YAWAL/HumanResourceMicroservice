package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/YAWAL/ERP-common-lib/database"
)

// Config structure represents general configuration for application
type Config struct {
	Port     string               `json:"port"`
	Database database.MongoConfig `json:"database"`
}

// ReadConfig performs reading config file, validates it and parse into Config structure
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
