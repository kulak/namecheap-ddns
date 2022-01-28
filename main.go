package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// remove time stamp from logger
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

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
	err = updateDnsHosts(ipStr, conf)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Updated hosts %+v.%s to IP: %v", conf.Hosts, conf.Domain, ipStr)
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

func updateDnsHosts(ipStr string, conf *Config) error {
	for _, eachHost := range conf.Hosts {
		if err := updateDnsHost(ipStr, conf, eachHost); err != nil {
			return err
		}
	}
	return nil
}

func updateDnsHost(ipStr string, conf *Config, host string) error {
	const format = "https://dynamicdns.park-your-domain.com/update?host=%v&domain=%v&password=%v&ip=%v"
	url := fmt.Sprintf(format, host, conf.Domain, conf.Password, ipStr)
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
	err = xmlDecode(content, &nr)
	if err != nil {
		log.Printf("Namecheap message: %v, url: %v", string(content), url)
		return fmt.Errorf("Failed to parse namecheap response: %v", err)
	}
	if nr.ErrCount != 0 {
		log.Printf("Namecheap message: %v, url: %v", string(content), url)
		return fmt.Errorf("Namecheap did not accept request: %v", nr.Errors.Err1)
	}
	// log.Printf("SUCCESS: %v", string(content))
	return nil
}

func xmlDecode(data []byte, v interface{}) error {
	d := xml.NewDecoder(bytes.NewReader(data))
	d.CharsetReader = identReader
	if err := d.Decode(v); err != nil {
		return fmt.Errorf("failed to xml decode: %v", err)
	}
	return nil
}

func identReader(encoding string, input io.Reader) (io.Reader, error) {
	return input, nil
}