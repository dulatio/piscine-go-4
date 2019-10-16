package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	var n32 int32
	a := [...]rune{48, 48, 48, 48, 48, 48, 48, 48, 48, 48}
	if n == 1 {
		n32 = 1
	} else if n == 2 {
		n32 = 2
	} else if n == 3 {
		n32 = 3
	} else if n == 3 {
		n32 = 3
	} else if n == 4 {
		n32 = 4
	} else if n == 5 {
		n32 = 5
	} else if n == 6 {
		n32 = 6
	} else if n == 7 {
		n32 = 7
	} else if n == 8 {
		n32 = 8
	} else {
		n32 = 9
	}
	var i int32
	for i = 0; i < n32; i++ {
		a[i] = i + 48
	}
	var index int32
	index = n32 - 1 
	for {
		for i = 0; i < n32; i++ {
			z01.PrintRune(a[i])
		}
		if a[0] == 48+10-n32 {
			break
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
		a[index]++
		if a[index] == (48+10-n32+index)+1 {		
			for ; index >= 0 && a[index] >= (48+10-n32+index); index-- {
			}
			a[index]++
			for index = index + 1; index < n32; index++ {
				a[index] = a[index-1] + 1
			}
			index = n32 - 1
		}
	}
	z01.PrintRune(10)
}
