package client

import (
	"io"
	"log"
	"net/http"
)

func GetSubscription(subUrl string) string {

	req, err := http.NewRequest("GET", subUrl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:143.0) Gecko/20100101 Firefox/143.0")

	cli := &http.Client{}
	resp, err := cli.Do(req)
	byte, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("read error", err)
	}
	base64Str := string(byte)
	log.Println("resp" + base64Str)

	return base64Str
}
