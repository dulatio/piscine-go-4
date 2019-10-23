package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printStr(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	counter := 0
	for i := range os.Args {
		if i != 0 {
			counter++
		}
	}
	if isEven(counter) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
