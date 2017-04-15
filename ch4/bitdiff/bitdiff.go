package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
)

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func byteArrToInt(xs []byte) int64 {
	i, _ := strconv.ParseInt(fmt.Sprintf("%x", xs), 16, 64)

	return i
}

func hamming(xs, ys [32]byte) int {
	returner := 0

	for i := 0; i < len(xs); i += len(xs) / 4 {
		xi := byteArrToInt(xs[i : i+len(xs)/4])
		yi := byteArrToInt(ys[i : i+len(ys)/4])

		xor := uint64(xi ^ yi)
		returner += PopCountByClearing(xor)
		//fmt.Printf("%x", xs[i:i+len(xs) / 4])

	}

	return returner
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: bitdiff <s1> <s2>")
		os.Exit(1)
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	h1 := sha256.Sum256([]byte(s1))
	h2 := sha256.Sum256([]byte(s2))

	fmt.Printf("%q\t%x\n", s1, h1)
	fmt.Printf("%q\t%x\n", s2, h2)

	fmt.Println("")

	fmt.Printf("Diff:\t%d", hamming(h1, h2))
}
