package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zoocasa/presignS3/lib/sign"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		signedURL, err := sign.URL(key)
		if err != nil {
			log.Print(err)
		}
		w.Write(signedURL)
	})

	fmt.Println("Server started")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
