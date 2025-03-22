package gandelbrot

import (
	"log"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	file, err := os.Create("render.png")
	if err != nil {
		log.Fatal(err)
	}

	Render(&RenderArgs{
		real:         -2.5,
		imaginary:    -2,
		complexWidth: 4.0,
		writer:       file,
	})

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
