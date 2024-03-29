package piscine

import "github.com/01-edu/z01"

func PrintNbrBase2(nbr int, base string) {
	len := StrLen(base)
	runes := []rune(base)
	valid := true
	if len < 2 {
		valid = false
	} else {
		for i := 0; i < len; i++ {
			if runes[i] == '-' || runes[i] == '+' {
				valid = false
			}
			for j := i + 1; j < len; j++ {
				if runes[i] == runes[j] {
					valid = false
				}
			}
		}
	}
	if !valid {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if nbr == 0 {
			z01.PrintRune(runes[0])
		} else {
			if nbr < 0 {
				z01.PrintRune('-')
				nbr *= -1
			}
			PrintNbrBaseRec(nbr, runes, len)
		}
	}
}