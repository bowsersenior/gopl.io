// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package mandelbrot

import (
	// "fmt"
	"image/color"
	_ "image/color/palette"
	"math/cmplx"
)

func MandelbrotColor(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// i := uint8(255 - contrast*n)
			return color.RGBA{R: (255 - contrast*n), G: (128 - contrast*n), B: (64 - contrast*n), A: 255}
			// i := int(contrast*n) % len(palette.Plan9)
			//
			// return palette.Plan9[i]
		}
	}
	return color.Black
}

func MandelbrotGray(z complex128) color.Color {
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

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func NewtonGray(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

func NewtonColor(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for n := uint8(0); n < iterations; n++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			alpha := 255 - contrast*n
			switch {
			case cmplx.Abs(z-1i) < 1e-2:
				return color.RGBA{R: 255, A: alpha}
			case cmplx.Abs(z-1) < 1e-2:
				return color.RGBA{G: 255, A: alpha}
			case cmplx.Abs(z+1) < 1e-2:
				return color.RGBA{B: 255, A: alpha}
			case cmplx.Abs(z+1i) < 1e-2:
				return color.RGBA{R: 124, B: 127, G: 64, A: alpha}
			}

			// return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
