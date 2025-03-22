// The gandelbrot package renders the Mandelbrot Set.
package gandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"slices"
	"sync"
	"time"
)

func calculateBlockRoot(renderWidth int) int {
	blockLength := 50

	for {
		if renderWidth%blockLength == 0 {
			break
		}

		blockLength += 1
	}

	return renderWidth / blockLength
}

func countIterations(
	real, imaginary float64,
	maximum int,
	maxOrbitLength int,
) int {
	count := 0
	cr := 0.0
	ci := 0.0

	// Record the history of cr and ci so we can bail early if we spot a loop.
	rHistory := make([]float64, maxOrbitLength)
	iHistory := make([]float64, maxOrbitLength)

	for {
		crr := cr * cr
		cii := ci * ci

		if crr+cii > 4 {
			return count
		}

		xt := crr - cii + real
		ci = 2*cr*ci + imaginary
		cr = xt

		// If cr and ci have both had these values together before then we know
		// we're in an infinite loop and it's time to bail.
		if hi := slices.Index(rHistory, cr); hi >= 0 && iHistory[hi] == ci {
			return maximum
		}

		count++

		if count >= maximum {
			return count
		}

		// We don't really care about the order of our history, so it's okay (and
		// quick) to overwrite older indices.
		rHistory[count%maxOrbitLength] = cr
		iHistory[count%maxOrbitLength] = ci
	}
}

func pixelToComplex(
	pixelPosition int,
	minPixelPosition int,
	blockWidth float64,
	minComplexPosition float64,
	maxComplexPosition float64,
) float64 {
	pc := float64(pixelPosition-minPixelPosition) / blockWidth
	return ((1.0 - pc) * minComplexPosition) + (pc * maxComplexPosition)
}

func worker(args *workerArgs) {
	for b := range args.blockChannel {
		// go-staticcheck warns that defers in this range loop won't run unless the
		// channel gets closed. And that's okay! I'll close the channel, I promise.
		defer args.waitGroup.Done()

		maxX := b.x + args.blockWidth
		maxY := b.y + args.blockWidth

		for x := b.x; x < maxX; x++ {
			real := pixelToComplex(
				x,
				b.x,
				args.blockWidthF,
				b.minReal,
				b.maxReal,
			)

			for y := b.y; y < maxY; y++ {
				imaginary := pixelToComplex(
					y,
					b.y,
					args.blockWidthF,
					b.minImaginary,
					b.maxImaginary,
				)

				count := countIterations(
					real,
					imaginary,
					args.renderArgs.maxIterations,
					args.renderArgs.maxOrbitLength,
				)

				colorValue := uint8((1 - (float64(count) / args.maxIterationsF)) * 255)

				args.img.Set(
					x,
					y,
					color.NRGBA{
						R: colorValue,
						G: colorValue,
						B: colorValue,
						A: 255,
					},
				)
			}
		}
	}
}

func Render(args *RenderArgs) {
	normalizeRenderArgs(args)

	start := time.Now()

	img := image.NewRGBA(
		image.Rect(0, 0, args.renderWidth, args.renderWidth),
	)

	blocksChannel := make(chan block)
	waitGroup := new(sync.WaitGroup)

	blockRoot := calculateBlockRoot(args.renderWidth)
	blockCount := blockRoot * blockRoot
	waitGroup.Add(blockCount)

	blockPixelWidth := args.renderWidth / blockRoot
	blockComplexWidth := args.complexWidth / float64(blockRoot)

	workerArgs := &workerArgs{
		blockChannel:   blocksChannel,
		blockWidth:     blockPixelWidth,
		blockWidthF:    float64(blockPixelWidth),
		img:            img,
		maxIterationsF: float64(args.maxIterations),
		renderArgs:     args,
		waitGroup:      waitGroup,
	}

	for range args.threadCount {
		go worker(workerArgs)
	}

	for blockX := range blockRoot {
		blockXF := float64(blockX)
		minReal := args.real + (blockComplexWidth * blockXF)
		maxReal := minReal + blockComplexWidth
		x := workerArgs.blockWidth * blockX

		for blockY := range blockRoot {
			blockYF := float64(blockY)

			minImaginary := args.imaginary + (blockComplexWidth * blockYF)
			maxImaginary := minImaginary + blockComplexWidth

			blocksChannel <- block{
				x:            x,
				y:            workerArgs.blockWidth * blockY,
				minReal:      minReal,
				maxReal:      maxReal,
				minImaginary: minImaginary,
				maxImaginary: maxImaginary,
			}
		}
	}

	// See! Didn't I promise you?
	close(blocksChannel)
	waitGroup.Wait()

	if err := png.Encode(args.writer, img); err != nil {
		log.Fatal(err)
	}

	log.Printf("rendered in %f seconds\n", time.Since(start).Seconds())
}
