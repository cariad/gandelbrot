package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
)

const (
	maxIterations = 800
	tileLength    = 400
)

func countIterations(px, py float64, maximum int) int {
	count := 0
	cx := 0.0
	cy := 0.0

	for {
		cxx := cx * cx
		cyy := cy * cy

		if cxx+cyy > 4 {
			return count
		}

		xt := cxx - cyy + px
		cy = 2*cx*cy + py
		cx = xt

		count += 1

		if count >= maximum {
			return count
		}
	}
}

func pixelToPlane(px, plMin, plMax float64) float64 {
	i := px / float64(tileLength)
	return ((1 - i) * plMin) + (i * plMax)
}

func renderTile(xMin, xMax, yMin, yMax float64, writer io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, tileLength, tileLength))

	for y := range tileLength {
		py := pixelToPlane(float64(y), yMin, yMax)

		for x := range tileLength {
			px := pixelToPlane(float64(x), xMin, xMax)
			count := countIterations(px, py, maxIterations)
			colorValue := uint8((1 - (float64(count) / maxIterations)) * 255)

			img.Set(x, y, color.NRGBA{
				R: colorValue,
				G: colorValue,
				B: colorValue,
				A: 255,
			})
		}
	}

	if err := png.Encode(writer, img); err != nil {
		log.Fatal(err)
	}
}
