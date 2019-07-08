package main

/* Exercise 8.5: Take an existing CPU-bound sequential program, such as the
Mandelbrot program of Section 3.3 or the 3-D surface computation of Section 3.2,
and execute its main loop in parallel using channels for communication. How much
faster does it run on a multiprocessor machine? What is the optimal number of
goroutines to use? */

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		concurrency            = 8
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	wg := sync.WaitGroup{}

	for c := 0; c < 2; c++ {
		offset := height / concurrency
		start := offset * c

		wg.Add(1)
		go func(first, offset int) {

			for py := first; py < first+offset; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}

			wg.Done()
		}(start, offset)
	}

	wg.Wait()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
