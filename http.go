package main

import (
	"log"
	"math"
	"net/http"
	"strconv"
)

type Tile struct {
	x int
	y int
	z int
}

func addHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "immutable, max-age=86400, no-transform, public")
}

func attachHttpHandlers() {
	http.Handle("/", handleRoot())
	http.HandleFunc("/tiles/{z}/{x}/{y}", tilesHandler)
}

func handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addHeaders(w)
		http.FileServer(http.Dir("frontend")).ServeHTTP(w, r)
	}
}

func readIntParam(r *http.Request, p string) int {
	s := r.PathValue(p)
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Printf("Failed to convert %s=%s to int\n", p, s)
	}

	return i
}

func readTile(r *http.Request) Tile {
	region := new(Tile)
	region.x = readIntParam(r, "x")
	region.y = readIntParam(r, "y")
	region.z = readIntParam(r, "z")
	return *region
}

func tilesHandler(w http.ResponseWriter, r *http.Request) {
	addHeaders(w)
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)

	region := readTile(r)
	planeLength := 4.0 / math.Pow(2, float64(region.z))
	xMin := (planeLength * float64(region.x)) - 2.0
	xMax := xMin + planeLength
	yMin := (planeLength * float64(region.y)) - 2.0
	yMax := yMin + planeLength

	renderTile(xMin, xMax, yMin, yMax, w)
}
