package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	a := [21]rune{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}

	for i := 20; n != 0; i-- {
		if n%10 == 0 {
			a[i] = '0'
		} else if n%10 == 1 {
			a[i] = '1'
		} else if n%10 == 2 {
			a[i] = '2'
		} else if n%10 == 3 {
			a[i] = '3'
		} else if n%10 == 4 {
			a[i] = '4'
		} else if n%10 == 5 {
			a[i] = '5'
		} else if n%10 == 6 {
			a[i] = '6'
		} else if n%10 == 7 {
			a[i] = '7'
		} else if n%10 == 8 {
			a[i] = '8'
		} else if n%10 == 9 {
			a[i] = '9'
		}
		n /= 10
	}

	index := 0
	for ; index < 21 && a[index] == '0'; index++ {

	}

	if index == 21 {
		index--
	}

	for ; index < 21; index++ {
		z01.PrintRune(a[index])
	}

	//z01.PrintRune(10)
}
