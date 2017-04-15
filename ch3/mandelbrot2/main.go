// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	m "gopl.io/ch3/mandelbrot2/mandelbrot"
	"gopl.io/ch3/surface/surface"

	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	xmin, ymin, xmax, ymax      = -2, -2, +2, +2
	defaultWidth, defaultHeight = 1024, 1024
)

type ImageParams struct {
	xmin, ymin, xmax, ymax, width, height int
}

func toComplex(px, py, height, width int) complex128 {
	y := float64(py)/float64(height)*(ymax-ymin) + ymin
	x := float64(px)/float64(width)*(xmax-xmin) + xmin
	z := complex(x, y)

	return z
}

func mandelbrotFromPixel(px, py, height, width int) color.Color {
	z := toComplex(px, py, height, width)
	return m.MandelbrotGray(z)
}

func avg(xs ...uint32) uint32 {
	returner := uint32(0)
	for _, x := range xs {
		returner += x
	}

	return (returner / uint32(len(xs)))
}

func mandelbrotColorFromPixelAvg(px, py, height, width int) color.Color {
	c1r, c1g, c1b, c1a := m.MandelbrotColor((toComplex(px, py, height, width))).RGBA()
	c2r, c2g, c2b, c2a := m.MandelbrotColor((toComplex(px+1, py, height, width))).RGBA()
	c3r, c3g, c3b, c3a := m.MandelbrotColor((toComplex(px, py+1, height, width))).RGBA()
	c4r, c4g, c4b, c4a := m.MandelbrotColor((toComplex(px+1, py+1, height, width))).RGBA()

	r := avg(c1r, c2r, c3r, c4r)
	g := avg(c1g, c2g, c3g, c4g)
	b := avg(c1b, c2b, c3b, c4b)
	a := avg(c1a, c2a, c3a, c4a)

	// fmt.Printf("%d - %d\t%d\t%d\t%d\n", r, c1r, c2r, c3r, c4r)

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

func mandelbrotGrayFromPixelAvg(px, py, height, width int) color.Color {
	c1r, c1g, c1b, c1a := m.MandelbrotGray((toComplex(px, py, height, width))).RGBA()
	c2r, c2g, c2b, c2a := m.MandelbrotGray((toComplex(px+1, py, height, width))).RGBA()
	c3r, c3g, c3b, c3a := m.MandelbrotGray((toComplex(px, py+1, height, width))).RGBA()
	c4r, c4g, c4b, c4a := m.MandelbrotGray((toComplex(px+1, py+1, height, width))).RGBA()

	r := avg(c1r, c2r, c3r, c4r)
	g := avg(c1g, c2g, c3g, c4g)
	b := avg(c1b, c2b, c3b, c4b)
	a := avg(c1a, c2a, c3a, c4a)

	// fmt.Printf("%d - %d\t%d\t%d\t%d\n", r, c1r, c2r, c3r, c4r)

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

func getFractalPNG(fn func(complex128) color.Color, out io.Writer, height int, width int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		// y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			// x := float64(px)/width*(xmax-xmin) + xmin
			// z := complex(x, y)
			// z := toComplex(px, py)
			// Image point (px, py) represents complex value z.
			// img.Set(px, py, mandelbrotGrayFromPixelAvg(px, py))

			img.Set(px, py, fn(toComplex(px, py, height, width)))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func IntFromForm(r *http.Request, k string, defaultVal int) int {
	var returner int

	if len(r.Form[k]) > 0 && r.Form[k][0] != "" {
		returner, _ = strconv.Atoi(r.Form[k][0])
		// fmt.Println(k+":", returner)
	} else {
		returner = defaultVal
	}

	return returner
}

func StringFromForm(r *http.Request, k string, defaultVal string) string {
	var returner string

	if len(r.Form[k]) > 0 && r.Form[k][0] != "" {
		returner = r.Form[k][0]
	} else {
		returner = defaultVal
	}

	return returner
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}

			imageType := StringFromForm(r, "image", "newton")
			height := IntFromForm(r, "height", defaultHeight)
			width := IntFromForm(r, "width", defaultWidth)

			switch imageType {
			case "newton":
				getFractalPNG(m.NewtonColor, w, height, width)
			case "mandelbrot":
				getFractalPNG(m.MandelbrotColor, w, height, width)
			case "surface":
				w.Header().Set("Content-type", "image/svg+xml")
				surface.Surface(w, height, width)
			}
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8081", nil))
		return
	} else {
		getFractalPNG(m.NewtonColor, os.Stdout, defaultHeight, defaultWidth)
	}
}
