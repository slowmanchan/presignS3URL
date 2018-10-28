package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/slowmanchan/presignS3URL/lib/sign"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		signedURL, err := sign.URL(key)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		w.Write(signedURL)
	})

	port := "9090"
	fmt.Printf("Server started on Port: %v...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
