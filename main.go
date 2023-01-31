package main

import (
	"fmt"
	"net/http"
	"log"
)

const addr = ":8087"

func main() {
	http.HandleFunc("/", handleIndex)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("index requested")
	fmt.Fprintf(w, "Hello from GitHub / Docker Hub")
}