package main

import (
	"fmt"
	"gopl.io/ch2/tempconv"
)

func main() {
	f := tempconv.Fahrenheit(100)

	fmt.Println(f)
	fmt.Println(tempconv.FToC(f))
	fmt.Println(tempconv.FToK(f))

	fmt.Println()

	c := tempconv.Celsius(37.778)

	fmt.Println(tempconv.CToF(c))
	fmt.Println(c)
	fmt.Println(tempconv.CToK(c))

	fmt.Println()

	k := tempconv.Kelvin(310.93)

	fmt.Println(tempconv.KToF(k))
	fmt.Println(tempconv.KToC(k))
	fmt.Println(k)
}
