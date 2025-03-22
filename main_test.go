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
		Real:         -2.5,
		Imaginary:    -2,
		ComplexWidth: 4.0,
		Writer:       file,
	})

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
