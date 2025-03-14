package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const baseZoom = 3.5

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

func iterationsToColour(count int, minCount int, countRange float64) uint8 {
	return uint8((1 - (float64(count-minCount) / countRange)) * 255)
}

func pixelToPlane(px int, pxMax float64, plMin float64, plMax float64) float64 {
	i := float64(px) / pxMax
	return ((1 - i) * plMin) + (i * plMax)
}

func render(width int, height int, zoom float64, filename string) {
	const maxIterations = 100
	const centerX = -0.5
	const centerY = 0.0

	widthF := float64(width)
	heightF := float64(height)

	ratio := heightF / widthF

	xWidth := baseZoom / zoom
	xMin := centerX - (xWidth / 2)
	xMax := centerX + (xWidth / 2)

	yHeight := xWidth * ratio
	yMin := centerY - (yHeight / 2)
	yMax := centerY + (yHeight / 2)

	counts := make([]int, width*height)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	minCount := maxIterations
	maxCount := 0

	for y := range height {
		indexOffset := y * width
		py := pixelToPlane(y, heightF, yMin, yMax)

		for x := range width {
			px := pixelToPlane(x, widthF, xMin, xMax)
			count := countIterations(px, py, maxIterations)
			counts[x+indexOffset] = count
			minCount = min(count, minCount)
			maxCount = max(count, maxCount)
		}
	}

	iterationsDelta := float64(maxCount - minCount)

	for y := range height {
		indexOffset := y * width

		for x := range width {
			value := iterationsToColour(
				counts[x+indexOffset],
				minCount,
				iterationsDelta,
			)

			img.Set(x, y, color.NRGBA{
				R: value,
				G: value,
				B: value,
				A: 255,
			})
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(file, img); err != nil {
		file.Close()
		log.Fatal(err)
	}
}
