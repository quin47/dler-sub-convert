package client

import (
	"io"
	"log"
	"net/http"
	"time"

	"gopkg.in/yaml.v3"
)

func GetSubscription(subUrl string) Subs {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, _ := client.Get(subUrl)
	byte, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("read error", err)
	}
	t := Subs{}
	yaml.Unmarshal(byte, &t)

	return t
}

type Subs struct {
	Proxies []Sub `yaml:"proxies,flow"`
}

type Sub struct {
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	TypeA    string `yaml:"type"`
	Password string `yaml:"password"`
	Udp      bool   `yaml:"udp"`
}
