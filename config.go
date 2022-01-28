package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config provides host specific data for Namecheap DDNS.
//
// Names are based on this URL example:
// https://dynamicdns.park-your-domain.com/update?host=[host]&domain=[domain_name]&password=[ddns_password]&ip=[your_ip]
type Config struct {
	Hosts    []string
	Domain   string
	Password string
}

func (c *Config) FromFile(fileName string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to read %s: %v", fileName, err)
	}
	if err := yaml.Unmarshal(content, c); err != nil {
		return fmt.Errorf("failed to decoded YAML %s file: %v", fileName, err)
	}
	return nil
}
