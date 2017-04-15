// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"

	"gopl.io/ch1/lissajous"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler

	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		var (
			cycles  = 5     // number of complete x oscillator revolutions
			res     = 0.001 // angular resolution
			size    = 100   // image canvas covers [-size..+size]
			nframes = 64    // number of animation frames
			delay   = 8     // delay between frames in 10ms units
		)

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		if len(r.Form["cycles"]) > 0 && r.Form["cycles"][0] != "" {
			cycles, _ = strconv.Atoi(r.Form["cycles"][0])
		}

		if len(r.Form["res"]) > 0 && r.Form["res"][0] != "" {
			res, _ = strconv.ParseFloat(r.Form["res"][0], 64)
		}

		if len(r.Form["size"]) > 0 && r.Form["size"][0] != "" {
			size, _ = strconv.Atoi(r.Form["size"][0])
		}

		if len(r.Form["nframes"]) > 0 && r.Form["nframes"][0] != "" {
			nframes, _ = strconv.Atoi(r.Form["nframes"][0])
		}

		if len(r.Form["delay"]) > 0 && r.Form["delay"][0] != "" {
			delay, _ = strconv.Atoi(r.Form["delay"][0])
		}

		lissajous.Draw(w, cycles, res, size, nframes, delay)
	})

	log.Fatal(http.ListenAndServe("localhost:8002", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//!-
