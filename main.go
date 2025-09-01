package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strconv"

	"github.com/quin47/dler-sub-convert/client"
)

func main() {
	fooHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query()["key"][0]
		subscription := client.GetSubscription("https://dler.cloud/subscribe/" + key + "??mu=trojan&lv=1%7C2%7C3")
		re := ""

		for _, s := range subscription.Proxies {
			re += "trojan://" + s.Password + "@" + s.Server + ":" + strconv.Itoa(s.Port) + "?allowInsecure=true#" + s.Name + "\n"
		}
		data := []byte(re)
		str := base64.StdEncoding.EncodeToString(data)
		w.Write([]byte(str))
		w.WriteHeader(http.StatusOK)

	})

	http.Handle("/subv/convert", fooHandler)
	log.Fatal(http.ListenAndServe(":9822", nil))

}
