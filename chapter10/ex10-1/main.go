package main

/*
Extend the jpeg program so that it converts any supported input format to any output format, using image.Decode to
detect the input format and a flag to select the output format.
*/

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

type Encoder interface {
	Encode(io.Writer, image.Image) error
}

type jpegEncoder struct {}
func (e *jpegEncoder) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, &jpeg.Options{Quality: 95})
}

type gifEncoder struct {}
func (g *gifEncoder) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, &gif.Options{NumColors: 256})
}

func main() {
	outFmt := flag.String("output", "jpeg", "output format")
	flag.Parse()

	var encoder Encoder
	switch *outFmt {
	case "jpeg":
		encoder = &jpegEncoder{}
	case "png":
		encoder = &png.Encoder{}
	case "gif":
		encoder = &gifEncoder{}
	default:
		fmt.Printf("unsupported output format: %s\n", *outFmt)
		os.Exit(1)
	}

	if err := convert(os.Stdin, os.Stdout, encoder); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer, encoder Encoder) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)


	return encoder.Encode(out, img)
}

