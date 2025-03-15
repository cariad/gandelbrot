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

func attachHttpHandlers() {
	http.Handle("/", http.FileServer(http.Dir("frontend")))
	http.HandleFunc("/tiles/{z}/{x}/{y}", tilesHandler)
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/png")

	region := readTile(r)
	planeLength := 4.0 / math.Pow(2, float64(region.z))
	xMin := (planeLength * float64(region.x)) - 2.0
	xMax := xMin + planeLength
	yMin := (planeLength * float64(region.y)) - 2.0
	yMax := yMin + planeLength

	renderTile(xMin, xMax, yMin, yMax, w)
}
