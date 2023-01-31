package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
)

const addr = ":8087"

func main() {
	http.HandleFunc("/", handleIndex)

	server := &http.Server{
        Addr: addr,
        ReadHeaderTimeout: 3 * time.Second,
    }

    err := server.ListenAndServe()

	if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("index requested")
	fmt.Fprintf(w, "Hello from GitHub / Docker Hub.")
}