package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	attachHttpHandlers()

	log.Printf("Listening on port %d…", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
