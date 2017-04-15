// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func anyNaN(x ...float64) bool {
	returner := false

	for _, f := range x {
		if math.IsNaN(f) {
			returner = true
			break
		}
	}

	return returner
}

func Surface(out io.Writer, height int, width int) {
	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
	zscale := float64(height) * 0.4         // pixels per z unit

	svg := ""

	diff := 0xff0000 - 0x0000ff
	unit := int(diff / 10)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, _ = corner(i, j, height, width, xyscale, zscale)
			_, _, _ = corner(i+1, j+1, height, width, xyscale, zscale)
		}
	}

	zrange := maxZ - minZ

	svg += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, height, width, xyscale, zscale)
			bx, by, bz := corner(i, j, height, width, xyscale, zscale)
			cx, cy, cz := corner(i, j+1, height, width, xyscale, zscale)
			dx, dy, dz := corner(i+1, j+1, height, width, xyscale, zscale)

			if anyNaN(ax, ay, bx, by, cx, cy, dx, dy) {
				//
			} else {
				z := (az + bz + cz + dz) / 4
				color := 0xff0000 - int((maxZ-z)/zrange*10)*unit

				svg += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: #%06x'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)

			}
		}
	}
	svg += fmt.Sprintln("</svg>")

	out.Write([]byte(svg))
}

var maxZ float64
var minZ float64

func corner(i int, j int, height int, width int, xyscale float64, zscale float64) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if z > maxZ {
		maxZ = z
	} else if z < minZ {
		minZ = z
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
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

//!-
