package main

import "github.com/01-edu/z01"

func main() {
	for c := 'a'; c <= 'z'; c++ {
		z01.PrintRune(rune(c))
	}
	r := 'a'
	s := string(r)
	_ = s
	z01.PrintRune('\n')
}
