package main

import (
	"log"
	"net/http"

	"github.com/quin47/dler-sub-convert/client"
)

func main() {
	fooHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query()["key"][0]
		base64Str := client.GetSubscription("https://dler.cloud/subscribe/" + key + "?mu=trojan&lv=1|2|3")
		w.Write([]byte(base64Str))
	})

	http.Handle("/subv/convert", fooHandler)
	log.Fatal(http.ListenAndServe(":9822", nil))

}
