package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	http.HandleFunc("/", rootHandler)

	log.Printf("Listening on port %d…", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
