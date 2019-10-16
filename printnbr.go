package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var a [20]int

	for i := 19; ; i-- {
		a[i] = n%10 + 48
		n /= 10
		if n == 0 {
			break
		}
	}

	index := 0
	for index := 0; a[index] == a[0]; index++ {

	}

	for ; index < 20; index++ {
		var r rune
		r = a[index]
		z01.PrintRune(r)
	}

	z01.PrintRune(10)
}
