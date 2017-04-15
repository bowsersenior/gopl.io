// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	KelvinDiffK   Kelvin  = 273.15
	AbsoluteZeroC Celsius = Celsius(-KelvinDiffK)
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%.5g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.5g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.5g°K", k) }

//!-
