package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	File = "conf.yaml"
)

type ws_attr struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type ws_server struct {
	Name   string    `yaml:"name"`
	Port   int       `yaml:"port"`
	Attr_l []string  `yaml:"attr_l"`
	Attr_m []ws_attr `yaml:"attr_m"`
}

type Conf struct {
	Ws ws_server `yaml:"ws-server"`
}

func (c *Conf) GetConf(file string) *Conf {

	yamlFile, err := ioutil.ReadFile(file)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
