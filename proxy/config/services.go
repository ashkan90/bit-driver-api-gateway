package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type ServiceConfig struct {
	Services []Service `yaml:"services"`
}

type Service struct {
	Name     string `yaml:"name,omitempty"`
	Target   string `yaml:"target,omitempty"`
	Strategy string `yaml:"strategy,omitempty"`
	Path     string `yaml:"path,omitempty"`
}

func (sc *ServiceConfig) ImportInto(fPath string) {
	var fName, _ = filepath.Abs(fPath)
	var yamlFile, err = ioutil.ReadFile(fName)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, sc)
}
