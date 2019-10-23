package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNbr(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNbr(points.y)
	z01.PrintRune('\n')
}

func printNbr(n int) {
	flag := false
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			n = 9223372036854775807
			flag = true
		} else {
			n = -n
		}
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
		if flag == true && index == 20 {
			z01.PrintRune('8')
		} else {
			z01.PrintRune(a[index])
		}
	}
}
