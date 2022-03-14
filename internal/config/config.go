package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Configs struct {
	Data         Database     `yaml:"database"`
	CloudServer  CloudServer  `yaml:"cloudServer"`
	CloudClient  CloudClient  `yaml:"cloudClient"`
	ObjectServer ObjectServer `yaml:"objectServer"`
	Location     Location     `yaml:"location"`
}

type Database struct {
	Mongo MongoDB `yaml:"mongodb"`
}

type ObjectServer struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	UseSSL          bool   `yaml:"useSSL"`
}

type MongoDB struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Ip           string `yaml:"ip"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"databaseName"`
}

type CloudServer struct {
	Platform  string `yaml:"platform"`
	LoginCode string `yaml:"loginCode"`
}

type CloudClient struct {
	DefaultStorage     float64 `yaml:"defaultStorage"`
	DefaultPhotoFolder string  `yaml:"defaultPhotoFolder"`
	DefaultFileFolder  string  `yaml:"defaultFileFolder"`
}

type Location struct {
	BingMapUrl string `yaml:"bingMapUrl"`
	BingMapKey string `yaml:"bingMapKey"`

	English            string `yaml:"English"`
	ChineseSimplified  string `yaml:"Chinese-Simplified"`
	ChineseTraditional string `yaml:"Chinese-Traditional"`
}

var Config *Configs

func InitConfig() {
	dir, err := filepath.Abs("../../test/config.yaml")
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
