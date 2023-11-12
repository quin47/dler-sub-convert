package client

import (
	"log"
	"net/http"
	"time"
	"io"
)

func GetSubscription(subUrl string) {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(subUrl)
	[]byte, error = io.ReadAll(resp.Body)

	yamlContent := string(byte)

	t := Sub{}
	err := yaml.Unmarshal([]byte(data), &t)

	log.Println(t)

}

type Sub struct {
	[]Proxies `yaml:"proxies"`
}

type  Proxies struct{
	server string
	port int
	name string
	type string
	password string
	udp bool
	skip-cert-verify bool

}