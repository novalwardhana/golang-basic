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
	}
}

var serverConfig *_Configuration

func init() {
	if serverConfig != nil {
		return
	}

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	configJSON, err := ioutil.ReadFile(filepath.Join(basePath, "conf", "config.json"))
	if err != nil {
		panic(err)
		return
	}

	serverConfig = new(_Configuration)
	err = json.Unmarshal(configJSON, &serverConfig)
	if err != nil {
		panic(err)
		return
	}
}

func Configuration() _Configuration {
	return *serverConfig
}
