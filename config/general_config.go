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

func (c *GeneralConfig) ImportInto(fPath string) {
	var fName, _ = filepath.Abs(fPath)
	var yamlFile, err = ioutil.ReadFile(fName)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, c)
}
