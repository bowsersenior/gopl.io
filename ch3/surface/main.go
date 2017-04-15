// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"gopl.io/ch3/surface/surface"
)

const (
	defaultWidth, defaultHeight = 600, 320 // canvas size in pixels
)

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

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "image/svg+xml")

			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}

			height := IntFromForm(r, "height", defaultHeight)
			width := IntFromForm(r, "width", defaultWidth)

			surface.Surface(w, height, width)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8081", nil))
		return
	} else {
		surface.Surface(os.Stdout, defaultHeight, defaultWidth)
	}
}

//!-
