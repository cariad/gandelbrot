package gandelbrot

import "io"

// Render arguments.
type RenderArgs struct {
	// The real (x) component of the top-left coordinate to render. Required.
	real float64

	// The imaginary (y) component of the top-left coordinate to render. Required.
	imaginary float64

	// The width of the square in the complex plane to render. Required.
	complexWidth float64

	// The writer to send the rendered image to.
	writer io.Writer

	// The maximum number of iterations to calculate for each point before
	// bailing. Omit or set <=0 for a sensible default.
	maxIterations int

	// Optional length of the calculation result stack for periodic orbit
	// detection. Set to <=0 for a sensible default.
	maxOrbitLength int

	// The width of the square bitmap to render. Omit or set <=0 for a sensible
	// default.
	renderWidth int

	// The number of worker threads perform calculations in. Omit or set <=0 for a
	// sensible default.
	threadCount int
}

func normalizeRenderArgs(args *RenderArgs) {
	if args.maxIterations <= 0 {
		args.maxIterations = 1_000
	}

	if args.maxOrbitLength <= 0 {
		args.maxOrbitLength = 50
	}

	if args.renderWidth <= 0 {
		args.renderWidth = 600
	}

	if args.threadCount < 1 {
		args.threadCount = 4
	}
}
