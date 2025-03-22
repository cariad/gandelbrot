package gandelbrot

import (
	"image"
	"sync"
)

// Worker thread initialisation arguments.
type workerArgs struct {
	blockChannel chan block
	blockWidth   int
	img          *image.RGBA
	renderArgs   *RenderArgs
	waitGroup    *sync.WaitGroup
}
