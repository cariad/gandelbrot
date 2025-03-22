package gandelbrot

// A region of a bitmap to render.
type block struct {
	// The X coordinate of the top-left pixel of the block within the image being
	// rendered.
	x int

	// The Y coordinate of the top-left pixel of the block within the image being
	// rendered.
	y int

	// The minimum real value of the complex region to render.
	minReal float64

	// The maxmimum real value of the complex region to render.
	maxReal float64

	// The minimum imaginary value of the complex region to render.
	minImaginary float64

	// The maximum imaginary value of the complex region to render.
	maxImaginary float64
}
