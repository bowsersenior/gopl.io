package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := []byte(scanner.Text())

		if len(os.Args) == 2 && os.Args[1] == "-384" {
			fmt.Printf("%x\n", sha512.Sum384(s))
		} else if len(os.Args) == 2 && os.Args[1] == "-512" {
			fmt.Printf("%x\n", sha512.Sum512(s))
		} else {
			fmt.Printf("%x\n", sha256.Sum256(s))
		}
	}
}
