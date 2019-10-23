package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	counter := 0
	for i := range os.Args {
		counter = i
	}
	for i := counter; i > 0; i-- {
		for _, r := range os.Args[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune(10)
	}
}
