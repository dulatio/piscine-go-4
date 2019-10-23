package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	for i, a := range os.Args {
		if i != 0 {
			for _, r := range a {
				z01.PrintRune(r)
			}
			z01.PrintRune(10)
		}
	}
}
