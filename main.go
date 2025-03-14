package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	// TODO: Stop mucking around.
	render(800, 600, 1.0, "render.png")

	attachHttpHandlers()

	log.Printf("Listening on port %d…", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
