package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	conf, err := loadConfig()
	if err != nil {
		log.Panic(err)
	}

	var ipStr string
	ipStr, err = getIP()
	if err != nil {
		log.Panic(err)
	}

	// update namecheap DDNS
	err = updateDns(ipStr, conf)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Updated %v to IP: %v", conf.Host, ipStr)
}

func getIP() (string, error) {
	// get IP address
	resp, err := http.Get("https://myip.supermicro.com/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var ipBytes []byte
	ipBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// strip \n character from response message
	rv := string(ipBytes)
	rv = strings.TrimSuffix(rv, "\n")
	return rv, nil
}

func loadConfig() (*Config, error) {
	// set configuration file name or use default
	var confFileName = "namecheap-ddns.yaml"
	if len(os.Args) == 2 {
		confFileName = os.Args[1]
	}

	// load configuration from yaml file
	conf := &Config{}
	err := conf.FromFile(confFileName)
	return conf, err
}

func updateDns(ipStr string, conf *Config) error {
	const format = "https://dynamicdns.park-your-domain.com/update?host=%v&domain=%v&password=%v&ip=%v"
	url := fmt.Sprintf(format, conf.Host, conf.Domain, conf.Password, ipStr)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to update IP: %v", err)
	}
	defer resp.Body.Close()
	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read namecheap response: %v", err)
	}
	var nr NamecheapResponse
	err = xml.Unmarshal(content, &nr)
	if err != nil {
		log.Printf("Namecheap message: %v, url: %v", string(content), url)
		return fmt.Errorf("Failed to parse namecheap response: %v", err)
	}
	if nr.ErrCount != 0 {
		log.Printf("Namecheap message: %v, url: %v", string(content), url)
		return fmt.Errorf("Namecheap did not accept request: %v", nr.Errors.Err1)
	}
	log.Printf("SUCCESS: %v", string(content))
	return nil
}
