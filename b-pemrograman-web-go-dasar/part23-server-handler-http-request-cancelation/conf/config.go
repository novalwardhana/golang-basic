package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type _Configuration struct {
	Server struct {
		Port         int           `json:"port"`
		ReadTimeout  time.Duration `json:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout"`
	} `json:"server"`
	Log struct {
		Verbose bool `json:"verbose"`
	} `json:"log"`
}

var configData *_Configuration

func init() {
	if configData != nil {
		return
	}

	basepath, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	configJSON, err := ioutil.ReadFile(filepath.Join(basepath, "conf", "config.json"))
	if err != nil {
		panic(err)
		return
	}

	configData = new(_Configuration)
	err = json.Unmarshal(configJSON, &configData)
	if err != nil {
		panic(err)
		return
	}
}

func Configuration() _Configuration {
	return *configData
}
