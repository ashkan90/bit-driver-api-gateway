package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type GeneralConfig struct {
	Server   Server    `yaml:"server"`
	Services []Service `yaml:"services"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Service struct {
	Name     string `yaml:"name,omitempty" json:"name,omitempty"`
	Target   string `yaml:"target,omitempty" json:"target,omitempty"`
	Strategy string `yaml:"strategy,omitempty" json:"strategy,omitempty"`
	Listen   string `json:"listen,omitempty"`
	Path     string `yaml:"path,omitempty" json:"path,omitempty"`
}

// NewConfig creates a new config struct with given yaml.
func NewConfig(fPath string) (GeneralConfig, error) {
	var conf GeneralConfig
	var fName, _ = filepath.Abs(fPath)
	var yamlFile, err = ioutil.ReadFile(fName)

	if err != nil {
		return GeneralConfig{}, err
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return GeneralConfig{}, err
	}

	return conf, nil
}
