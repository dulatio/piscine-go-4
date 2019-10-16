package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	nn = int32(n)
	if nn < 0 {
		z01.PrintRune('-')
		nn = -nn
	}
	var a [20]int32

	for i := 19; ; i-- {
		a[i] = nn%10 + 48
		nn /= 10
		if nn == 0 {
			break
		}
	}

	index := 0
	for index := 0; a[index] == a[0]; index++ {

	}

	for ; index < 20; index++ {
		z01.PrintRune(a[index])
	}

	z01.PrintRune(10)
}
