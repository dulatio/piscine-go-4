package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	for _, r := range os.Args[0] {
		z01.PrintRune(r)
	}
	z01.PrintRune(10)
}
