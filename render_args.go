package gandelbrot

import "io"

// Render arguments.
type RenderArgs struct {
	// The real (x) component of the top-left coordinate to render. Required.
	Real float64

	// The imaginary (y) component of the top-left coordinate to render. Required.
	Imaginary float64

	// The width of the square in the complex plane to render. Required.
	ComplexWidth float64

	// The writer to send the rendered image to.
	Writer io.Writer

	// The maximum number of iterations to calculate for each point before
	// bailing. Omit or set <=0 for a sensible default.
	MaxIterations int

	// Optional length of the calculation result stack for periodic orbit
	// detection. Set to <=0 for a sensible default.
	MaxOrbitLength int

	// The width of the square bitmap to render. Omit or set <=0 for a sensible
	// default.
	RenderWidth int

	// The number of worker threads to perform calculations in. Omit or set <=0
	// for a sensible default.
	ThreadCount int
}

func normalizeRenderArgs(args *RenderArgs) {
	if args.MaxIterations <= 0 {
		args.MaxIterations = 1_000
	}

	if args.MaxOrbitLength <= 0 {
		args.MaxOrbitLength = 50
	}

	if args.RenderWidth <= 0 {
		args.RenderWidth = 600
	}

	if args.ThreadCount < 1 {
		args.ThreadCount = 4
	}
}
