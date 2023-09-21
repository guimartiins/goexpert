package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request started")
	defer log.Println("request finished")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request done with success")
		w.Write([]byte("Request done with success"))
	case <-ctx.Done():
		log.Println("Request cancelled by client")
		http.Error(w, "Request cancelled by client", http.StatusRequestTimeout)
	}
}
