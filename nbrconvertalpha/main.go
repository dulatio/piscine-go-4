package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	upper := false
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		if i == 1 && v == "--upper" {
			upper = true
			continue
		}
		if strToNbr(v) >= 1 && strToNbr(v) <= 26 {
			if upper {
				z01.PrintRune(rune(strToNbr(v) + 'A' - 1))
			} else {
				z01.PrintRune(rune(strToNbr(v) + 'a' - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
}

func strToNbr(s string) int {
	res := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			res *= 10
			res += int(r - '0')
		} else {
			return -1
		}
	}
	return res
}
