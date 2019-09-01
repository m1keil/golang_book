// Supersampling is a technique to reduce the effect of pixelation by computing the color value at several points within
// each pixel and taking the average. The simplest method is to divide each pixel into four “subpixels.” Implement it.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, supersample(img)) // NOTE: ignoring errors
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

func supersample(img *image.RGBA) image.Image {
	width := img.Rect.Dx() / 2
	height := img.Rect.Dy() / 2

	out := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			x1 := img.RGBAAt(px*2, py*2)
			x2 := img.RGBAAt(px*2+1, py*2)
			y1 := img.RGBAAt(px*2, py*2+1)
			y2 := img.RGBAAt(px*2+1, py*2+1)

			var Ravg, Gavg, Bavg, Aavg uint32
			Ravg = (uint32(x1.R) + uint32(x2.R) + uint32(y1.R) + uint32(y2.R)) / 4
			Gavg = (uint32(x1.G) + uint32(x2.G) + uint32(y1.G) + uint32(y2.G)) / 4
			Bavg = (uint32(x1.B) + uint32(x2.B) + uint32(y1.B) + uint32(y2.B)) / 4
			Aavg = (uint32(x1.A) + uint32(x2.A) + uint32(y1.A) + uint32(y2.A)) / 4

			z := color.RGBA{
				R: uint8(Ravg),
				G: uint8(Gavg),
				B: uint8(Bavg),
				A: uint8(Aavg),
			}
			out.Set(px, py, z)
		}
	}

	return out
}
