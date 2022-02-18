package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Configs struct {
	Data Database `yaml:"database"`
}

type Database struct {
	Mongo MongoDB `yaml:"mongodb"`
}

type MongoDB struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Ip           string `yaml:"ip"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"databaseName"`
}

var Config *Configs

func InitConfig() {
	dir, err := filepath.Abs("../../configs/config.yaml")
	if err != nil {
		return
	}

	file, err := ioutil.ReadFile(dir)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		return
	}
}
