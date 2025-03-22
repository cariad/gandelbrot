package gandelbrot

import (
	"image"
	"sync"
)

// Worker thread initialisation arguments.
type workerArgs struct {
	blockChannel chan block
	blockWidth   int

	// The float64 value of blockWidth. This is precalculated to save a million
	// casts at runtime.
	blockWidthF float64

	img *image.RGBA

	// The float64 value of the maximum number of iterations to perform. This is
	// precalculated to save a million casts at runtime.
	maxIterationsF float64

	renderArgs *RenderArgs
	waitGroup  *sync.WaitGroup
}
